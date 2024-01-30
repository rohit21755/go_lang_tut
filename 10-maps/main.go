package main

import "fmt"

func main() {
	fmt.Println("Hello maps")

	lang := make(map[string]string)
	lang["JS"] = "javascript"
	lang["PY"] = "python"
	fmt.Println(lang)
	// deleting
	// delete(lang, "JS")

	// looping
	// we can also use comma okay syntax for key and value
	for key, value := range lang {
		fmt.Printf("For key %v, value is %v", key, value)
	}

}
