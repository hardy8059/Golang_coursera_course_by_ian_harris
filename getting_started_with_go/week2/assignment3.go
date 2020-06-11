package main

import (
	"fmt"
	"strings"
)

func find_chars() {
	var x string
	fmt.Println("Please Enter the string: ")
	fmt.Scan(&x)
	x = strings.ToLower(x)
	i := strings.Index(x, "i")
	n := strings.Index(x, "n")
	a := strings.Index(x, "a")
	switch {
	case i != 0:
		fmt.Printf("Not Found!")
	case n != len(x)-1:
		fmt.Printf("Not Found!")
	case a == -1:
		fmt.Printf("Not Found!")
	default:
		fmt.Printf("Found!")
	}
}

func main() {
	find_chars()
}
