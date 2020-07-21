package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func inputNumbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func sortIntegers(intList []int, c chan []int, wg *sync.WaitGroup) {
	fmt.Println(intList)
	sort.Ints(intList)
	c <- intList
	wg.Done()
}

func main() {
	var series []int
	var wg sync.WaitGroup
	c := make(chan []int, 4)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	series = inputNumbers(scanner.Text())
	seriesLength := int(math.Round(float64(len(series) / 4)))

	go sortIntegers(series[0:seriesLength], c, &wg)
	go sortIntegers(series[seriesLength:2*seriesLength], c, &wg)
	go sortIntegers(series[2*seriesLength:3*seriesLength], c, &wg)
	go sortIntegers(series[3*seriesLength:], c, &wg)
	wg.Add(4)
	wg.Wait()
	fmt.Println("Execution over")
	s1 := <-c
	s2 := <-c
	mergedList := append(s1, s2...)
	s3 := <-c
	mergedList = append(mergedList, s3...)
	s4 := <-c
	mergedList = append(mergedList, s4...)
	fmt.Println("Sorted list is")
	wg.Add(1)
	go sortIntegers(mergedList, c, &wg)
	fmt.Println(<-c)
}
