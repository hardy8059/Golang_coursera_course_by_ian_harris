package main

import (
	"encoding/json"
	"fmt"
)

func printJson() {
	var input string
	nd := map[string]string{
		"name":    "",
		"address": "",
	}
	fmt.Println("Enter the name: ")
	fmt.Scan(&input)
	nd["name"] = input
	fmt.Println("Enter the address: ")
	fmt.Scan(&input)
	nd["address"] = input
	ndJson, _ := json.Marshal(nd)
	fmt.Println(string(ndJson))
}

func main() {
	printJson()
}
