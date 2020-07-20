package main

import (
	"fmt"
	"reflect"
)

func inputSeries() []int {
	var input int
	var series []int
	fmt.Println("Please enter the series of ten digits below:")
	for inputCount := 0; inputCount < 10; inputCount++ {
		fmt.Scan(&input)
		fmt.Println(input)
		series = append(series, input)
	}
	return series
}

func swapIntegers(s []int, pos int) {
	swapObject := reflect.Swapper(s)
	swapObject(pos, pos+1)
}

func BubbleSort(series []int) {
	for i := 0; i < len(series)-1; i++ {
		for j := 0; j < len(series)-1-i; j++ {
			if series[j] > series[j+1] {
				swapIntegers(series, j)
			}
		}
	}
}

func main() {
	s := inputSeries()
	BubbleSort(s)
	fmt.Println(s)
}
