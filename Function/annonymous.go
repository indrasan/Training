package main

import "fmt"

//annonymous function one usecase is to pass function as parameter in function
func main() {

	fmt.Println("Welcome to annonymous function")
	func(msg string) {

		fmt.Println("message is :", msg)
	}("Good Morning")

	//("Good Evening")

	fun := func(a, b int) int {

		return a + b
	}
	fmt.Println(fun(5, 10))
}
