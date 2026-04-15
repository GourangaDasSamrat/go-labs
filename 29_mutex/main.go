package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup

	newPost := post{views: 0}

	for range 200 {
		wg.Add(1)
		newPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(newPost.views)
}
