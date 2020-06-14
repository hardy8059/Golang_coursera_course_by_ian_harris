package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strings"
)

func readNameFile() {
	type Name struct {
		fname string
		lname string
	}
	var filename string
	fmt.Println("Please Enter the file name: ")
	fmt.Scan(&filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Encountered Error %s", err)
	}
	defer f.Close()

	n := []Name{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")
		n = append(n, Name{fname: res[0], lname: res[1]})
	}
	for i, p := range n {
		fmt.Printf("Name %d's First name is %s and Last name is %s \n", i+1, p.fname, p.lname)
	}
}

func main() {
	readNameFile()
}
