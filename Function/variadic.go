package main

import "fmt"

func add(a int, b int, adds ...int) int { //here a & b passd to make it mandatory to pass atleast 2 value
	var sum int = a + b

	for _, v := range adds {
		sum += v
	}
	return sum

}

func main() {

	sum := add(1, 2, 3, 4, 5)
	fmt.Println(sum)
}
