package main

import (
	"fmt"
)


func Swap(sli []int, index int) {

	if index > len(sli) {
		fmt.Println("Error : ArrayIndexOutOfBound")
	} else {
		sli[index + 1], sli[index] = sli[index], sli[index + 1]
	}
}


func BubbleSort(list []int) {

	for i := len(list); i >= 2; i--{
		for j := 0; j < i-1; j++{
			if list[j] > list[j + 1]{
				Swap(list, j)
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
	for _, value:= range integerList {
         fmt.Printf("%v ", value)
    }
}