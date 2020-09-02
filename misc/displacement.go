  
package main

import (
	"fmt"
	"math"
)


func GenDisplaceFn(a, v0, s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return (0.5*a*math.Pow(t, 2) + v0*t + s0)
	}
}


func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter acceleration, initial velocity and displacement separated by a space:  ")
	fmt.Scanln(&a, &v0, &s0)
	fmt.Print("Enter time:")
	fmt.Scanln(&t)
	
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}

