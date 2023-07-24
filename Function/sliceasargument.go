package main

import "fmt"

func slicearg(slice []int) {

	slice = append(slice, 100)

	fmt.Println(slice)
}

func main() {
	slc := []int{1, 2, 3, 4, 5}
	slicearg(slc)
	fmt.Println(slc)
}
