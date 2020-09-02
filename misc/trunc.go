package main

import (
	"fmt"
	"math"
)

func main() {
	var fp float64

	fmt.Println("Enter floating point number: ")
	fmt.Scanln(&fp)
	fmt.Println(int64(math.Floor(fp)))
}