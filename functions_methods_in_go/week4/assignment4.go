package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (c Cow) Eat()   { fmt.Println("Grass") }
func (c Cow) Move()  { fmt.Println("Walk") }
func (c Cow) Speak() { fmt.Println("Moo") }

func (b Bird) Eat()   { fmt.Println("Worms") }
func (b Bird) Move()  { fmt.Println("Fly") }
func (b Bird) Speak() { fmt.Println("Peep") }

func (s Snake) Eat()   { fmt.Println("Mice") }
func (s Snake) Move()  { fmt.Println("Slither") }
func (s Snake) Speak() { fmt.Println("Hsss") }

func printDetail(a Animal, detail string) {
	switch detail {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	var animal Animal
	var c Cow
	var b Bird
	var s Snake
	var command, animalName, detail string
	for {
		fmt.Print(">")
		fmt.Scanf("%s %s %s ", &command, &animalName, &detail)
		if command == "newanimal" {
			switch detail {
			case "cow":
				c.name = animalName
			case "bird":
				b.name = animalName
			case "Snake":
				s.name = animalName
			}
			fmt.Println("Created it!")
		}
		if command == "query" {
			switch animalName {
			case c.name:
				animal = c
			case b.name:
				animal = b
			case s.name:
				animal = s
			}
			printDetail(animal, detail)
		}
	}
}
