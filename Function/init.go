package main

import "fmt"

//init is first function to be called and its used to initialize global variable.after that main is called
// its called by system .

func init() {
	fmt.Println("Begin Init") //init order is in order of its defination.
}
func init() {
	fmt.Println("Begin Init1")
}

func main() {

	fmt.Println("I am main")
}
