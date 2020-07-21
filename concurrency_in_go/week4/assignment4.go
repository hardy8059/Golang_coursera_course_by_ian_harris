package main

import (
	"fmt"
	"sync"
)

type chopS struct {
	sync.Mutex
}

type Philosopher struct {
	lChop, rChop  *chopS
	num           int
	eatingCounter int
}

func initialise() ([]*chopS, []*Philosopher) {
	// Initialise Chopsticks
	cSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}

	// Initialise Philosophers
	p := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		p[i] = &Philosopher{lChop: cSticks[i], rChop: cSticks[(i+1)%5], num: i + 1}
	}
	return cSticks, p
}

func takePermission(request chan bool, permission chan bool) {
	i := 1
	for x := range request {
		i += 1
		fmt.Println("Currently Eating: ", len(request), i)
		permission <- x
	}
}

func (p Philosopher) eat(request chan bool, permission chan bool, wg *sync.WaitGroup) {
	request <- true
	requestState := <-permission
	if requestState {
		if p.num%2 == 0 {
			fmt.Println("Starting to eat ", p.num)
			p.lChop.Lock()
			p.rChop.Lock()
			fmt.Println("Finishing eating ", p.num)
			p.rChop.Unlock()
			p.lChop.Unlock()
			wg.Done()
		} else {
			fmt.Println("Starting to eat ", p.num)
			p.rChop.Lock()
			p.lChop.Lock()
			fmt.Println("Finishing eating ", p.num)
			p.lChop.Unlock()
			p.rChop.Unlock()
			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	request := make(chan bool, 2)
	permission := make(chan bool, 2)
	_, ph := initialise()
	fmt.Println(ph)
	wg.Add(15)
	go takePermission(request, permission)
	for i := 0; i < 15; i++ {
		j := i % 5
		go ph[j].eat(request, permission, &wg)
	}
	wg.Wait()
	fmt.Println("Execution Complete")
}
