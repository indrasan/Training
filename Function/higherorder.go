package main

//As function which takes function as argument is called higher order function. for example iterate function.
import "fmt"

func iterate(slicestr []string, fn func(string)) {

	for _, val := range slicestr {
		fn(val)
	}
}

func main() {

	slicestr := []string{"How", "are", "you"}
	fn := func(msg string) {
		fmt.Println(msg)
	}

	iterate(slicestr, fn)
}
