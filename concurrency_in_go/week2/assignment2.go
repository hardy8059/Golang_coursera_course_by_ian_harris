package main

import (
	"fmt"
)

var x = 1

func thread1() {
	fmt.Println(x)
}

func thread2() {
	x = x + 1
}

func main() {
	go thread1()
	go thread2()
	//time.Sleep(5 * time.Millisecond) // Not a recommended approach but included
	// as it works in this case. Uncomment to see race condition.
}

//Suppose the interleaving order executes the printf statement
// from thread 1 before the x = x+1 of thread 2, it will result in 2,
// whereas if thread 2 is executed first and then thread 1 it will result in
// 3. Hence the output is varying as per interleaving order which is the race condition.
