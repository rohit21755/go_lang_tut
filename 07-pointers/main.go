package main

import "fmt"

func main() {
	// var one *int
	two := 6
	var ptr = &two
	fmt.Println(ptr)
	fmt.Println(*ptr)

}
