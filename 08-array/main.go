package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "kiwi"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var vg = [5]string{"potato", "beans", "mush"}
	fmt.Println(vg)
	fmt.Println(len(vg))
}
