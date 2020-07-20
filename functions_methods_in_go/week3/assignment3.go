package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat(animal string) string {
	switch animal {
	case "cow":
		a.food = "Grass"
	case "bird":
		a.food = "Worms"
	case "snake":
		a.food = "Mice"
	}
	return a.food
}

func (a *Animal) Move(animal string) string {
	switch animal {
	case "cow":
		a.locomotion = "Walk"
	case "bird":
		a.locomotion = "Fly"
	case "snake":
		a.locomotion = "Slither"
	}
	return a.locomotion
}
func (a *Animal) Speak(animal string) string {
	switch animal {
	case "cow":
		a.noise = "Moo"
	case "bird":
		a.noise = "Peep"
	case "snake":
		a.noise = "Hsss"
	}
	return a.noise
}

func getDetail(animal string, detail string) string {
	var a Animal
	var result string
	switch detail {
	case "eat":
		result = a.Eat(animal)
	case "move":
		result = a.Move(animal)
	case "speak":
		result = a.Speak(animal)
	}
	return result
}

func main() {
	var animal, detail string
	for {
		fmt.Print(">")
		fmt.Scanf("%s %s ", &animal, &detail)
		fmt.Println(getDetail(animal, detail))
	}
}
