package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))
	createDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
