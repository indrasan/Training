package main

import "fmt"

func circleArea(radius float32) (area float32, peri float32) {
	area = 3.4195 * float32(radius) * float32(radius)
	peri = 2 * 3.4195 * float32(radius)

	return //no need to put return variable name as already its defined return variable name in signature

}

func main() {

	fmt.Println(circleArea(1.2))
}
