package main

import "fmt"

func main() {
	fmt.Println("hello struct")
	// no inheritance, no super class and parent
	rohit := User{"rohit", "cc@gmail.com", true, 21}
	fmt.Printf("%+v", rohit)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
