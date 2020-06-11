package main

import (
	"fmt"
)

func float_to_int() (int){
	var num float64
	fmt.Println("Please Enter a floating point digit: ")
	fmt.Scan(&num)
	y := int(num)
	return y
}

func main(){
	fmt.Print(float_to_int())
}