package main

import "fmt"

func messageDeferred(msg string) {

	defer fmt.Println("Inside defer", msg) // defer is pushed in stack in sequence , printed in LIFO Order
	fmt.Println("Hello")                   //2,4

}

func main() {
	fmt.Println("Demo of Defer:")         //1
	defer messageDeferred("Good Morning") //4
	defer messageDeferred("Good Evening") //3

}
