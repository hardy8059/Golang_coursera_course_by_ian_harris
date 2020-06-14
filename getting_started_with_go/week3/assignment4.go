package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sort_Integers() {
	sl := make([]int, 0, 3)
	fmt.Println(cap(sl))
	var x string
	for {
		fmt.Println("Please enter integer: ")
		fmt.Scan(&x)
		y, _ := strconv.Atoi(x)
		if x != "X" {
			sl = append(sl, y)
		} else {
			break
		}
		sort.Ints(sl)
		fmt.Printf("Sorted slice is: %d\n", sl)
	}
}

func main() {
	sort_Integers()
}
