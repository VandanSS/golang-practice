package main

import (
	"fmt"
)

func BubbleSort(list []int) {

	for i := len(list); i >= 2; i-- {
		for j := 0; j < i-1; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
	}
}

func main() {

	var integerList []int
	var num int

	fmt.Println("Enter a sequence of 10 integers:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		integerList = append(integerList, num)
	}

	BubbleSort(integerList)
	for _, value := range integerList {
		fmt.Printf("%v ", value)
	}
}
