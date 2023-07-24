package main

import "fmt"

/*  function syntax
func functionname(argument1 datatype,argument2 datatype) return_datatype
{
	body
}
*/
func add(a, b int) int {

	return a + b
}

func divide(a, b int) (int, bool) {

	if b == 0 {
		return 0, false
	} else {

		val := a / b
		return val, true

	}

}
func circleArea(radius float32) (float32, float32) {
	area := 3.4195 * float32(radius) * float32(radius)
	peri := 2 * 3.4195 * float32(radius)

	return area, peri

}

func main() {

	sum := add(5, 10)
	fmt.Println(sum)

	sum = add(add(10, 20), add(30, 40))
	fmt.Println(sum)

	val, eval := divide(100, 10)
	if eval == true {
		fmt.Println(val)
	} else {

		fmt.Println("exception")
	}

	val, eval = divide(100, 0)
	if eval == true {
		fmt.Println(val)
	} else {

		fmt.Println("exception")
	}

	area, peri := circleArea(2.14)
	fmt.Println(area, peri)

}
