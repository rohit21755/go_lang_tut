package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("goodd ")
	input, _ := reader.ReadString('\n')
	// for error thing we can use the syntax like input, err = ......
	// its work like try and catch
	//  if its about the error only then _, err = ......
	fmt.Println(input)

}
