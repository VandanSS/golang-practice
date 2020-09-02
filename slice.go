package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var input string
	pass := 1
	slc := make([]int, 0, 3)
	for {
		fmt.Printf("Pass %d \n",pass)
		fmt.Scanln(&input)	
		if input == "X" {
			break
		} else {
			input, err := strconv.Atoi(input)
			if err == nil {
				slc = append(slc, input)
				sort.Ints(slc)
				fmt.Printf("Slice : %v \n", slc)
			}
		}
		pass++
	}
}