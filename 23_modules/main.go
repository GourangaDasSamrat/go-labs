package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define structure to hold the post data from JSONPlaceholder API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/", handleServeHome).Methods("GET")

	// Start the server
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// handleServeHome fetches posts from JSONPlaceholder and serves them dynamically
func handleServeHome(w http.ResponseWriter, r *http.Request) {
	// Fetch posts from JSONPlaceholder API
	posts, err := fetchPosts()
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	// Define template data with the fetched posts
	data := struct {
		Title string
		Posts []Post
	}{
		Title: "Posts from JSONPlaceholder",
		Posts: posts,
	}

	// Parse and render the HTML template
	tmpl, err := template.New("home").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					padding: 20px;
					background-color: #f4f4f4;
				}
				h1 {
					color: #333;
				}
				.post {
					background-color: #fff;
					padding: 20px;
					border-radius: 8px;
					box-shadow: 0 0 10px rgba(0,0,0,0.1);
					margin-bottom: 20px;
				}
				.post h3 {
					color: #4CAF50;
				}
			</style>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			{{range .Posts}}
				<div class="post">
					<h3>{{.Title}}</h3>
					<p>{{.Body}}</p>
				</div>
			{{end}}
		</body>
		</html>
	`)

	// Error handling for template parsing
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Unable to load the page", http.StatusInternalServerError)
		return
	}

	// Execute template with data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// fetchPosts fetches a list of posts from the JSONPlaceholder API
func fetchPosts() ([]Post, error) {
	// JSONPlaceholder API URL for posts
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body using io.ReadAll (modern Go)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the JSON response
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return posts, nil
}
