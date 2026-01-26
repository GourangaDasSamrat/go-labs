package main

import "fmt"

// constant
const MyName string = "Gouranga Das Samrat" // this is constant like other language or math, means its can't be reassign or change
// also here first letter is capital letter means its a `public` variable, so we ca use this in anywhere under that package

func main() {
	// string
	var coffeeName string = "Mocha"
	fmt.Println(coffeeName)
	fmt.Printf("Variable type is: %T \n", coffeeName)

	// integer
	var servings int = 10
	fmt.Println(servings)
	fmt.Printf("Variable type is: %T \n", servings)

	// boolean
	var isInStock bool = true
	fmt.Println(isInStock)
	fmt.Printf("Variable type is: %T \n", isInStock)

	// untyped integer
	var cupSize uint8 = 255 // this is 8 bit means it can store 0 to 255
	fmt.Println(cupSize)
	fmt.Printf("Variable type is: %T \n", cupSize)

	// float
	var price float32 = 255.2137837483748347834798409 // this is 32 bit means it can store 5 values after decimal
	fmt.Println(price)
	fmt.Printf("Variable type is: %T \n", price)

	var latitude float64 = 255.2137837483748347834798409 // this is 64 bit means it can store 14 values after decimal
	fmt.Println(latitude)
	fmt.Printf("Variable type is: %T \n", latitude)

	//  default values
	var defValue int // default value is 0
	fmt.Println(defValue)
	fmt.Printf("Variable type is: %T \n", defValue)

	var defValue2 string // default value is a blank new line
	fmt.Println(defValue2)
	fmt.Printf("Variable type is: %T \n", defValue2)

	// aliases
	var aliasOne byte = 10 // alias for uint8
	fmt.Println(aliasOne)
	fmt.Printf("Variable type is: %T \n", aliasOne)

	// implicit types
	var mySite = "https:/gouranga.qzz.io" // lexer is automatically cast string type here
	fmt.Println(mySite)
	fmt.Printf("Variable type is: %T \n", mySite)

	// no var style
	myGitHub := "https://github.com/GourangaDasSamrat" // this is short variable declaration operator, its allowed to declare inside any method in this case main method but we can't declare it outside the method
	fmt.Println(myGitHub)
	fmt.Printf("Variable type is: %T \n", myGitHub)

	fmt.Println(MyName)
}
