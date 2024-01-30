package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")

	var fruits = []string{}
	fmt.Println("Types of fruitslist is %T", fruits)

	fruits = append(fruits, "mango", "banana")
	fmt.Println(fruits)

	// coloumn syntax

	fruits = append(fruits[1:]) // it remove the first element of the list
	// [1:3] start with 1 and stops at 3(not included)

	highScores := make([]int, 4)
	highScores[0] = 234

	highScores[2] = 676
	highScores[3] = 656
	// if you add one more it show erro out of boud
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 434)
	// it will reallocate the memory and all values will be acomadated
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
