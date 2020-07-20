package main

import (
	"fmt"
	"math"
)

type displacementVariables struct {
	acc                 float64
	initialVelocity     float64
	initialDisplacement float64
}

func GenDisplaceFn(variables displacementVariables) func(dispTime float64) float64 {
	return func(dispTime float64) float64 {
		return 0.5*variables.acc*math.Pow(dispTime, 2) + variables.initialVelocity*dispTime + variables.initialDisplacement
	}
}

func main() {
	var d displacementVariables
	fmt.Println("Enter acceleration value:")
	fmt.Scan(&d.acc)
	fmt.Println("Enter initial velocity value:")
	fmt.Scan(&d.initialVelocity)
	fmt.Println("Enter initial displacement value:")
	fmt.Scan(&d.initialDisplacement)
	fn := GenDisplaceFn(d)
	var t float64
	fmt.Println("Enter time value:")
	fmt.Scan(&t)
	fmt.Println("Displacement value is: ", fn(t))
}
