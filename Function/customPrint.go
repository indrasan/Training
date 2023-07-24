package main

import (
	"fmt"
	"strconv"
)

func mycustomPrint(format string, inputs ...interface{}) { //first input is string to make mandatory to keep first value string.
	fmt.Println(format, inputs) //unformatted string.

	s := format
	for _, val := range inputs {
		switch val.(type) {
		case int:
			s += " " + strconv.Itoa(val.(int))
		case string:
			s += " " + val.(string)

		}
	}
	fmt.Println(s) // formatted string
}

func main() {
	mycustomPrint("Hello", 1, "How are You", 123, 10.5)
}
