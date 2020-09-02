package main

import (
	"fmt"
	"sort"
)

func sortArr(list []int, c chan []int) {
	fmt.Println("Sorting subarray: ", list)
	sort.Ints(list)	
	c <- list
}

func main() {

	var integerList, slc1, slc2, slc3, slc4 []int
	var l1, l2, l3 int
	var length, num int
	
	fmt.Println("Enter number of integers you want to sort:")
	fmt.Scan(&length)
	
	fmt.Println("Enter integer list")
	
	for i := 0; i < length; i++ {
		fmt.Scan(&num)
		integerList = append(integerList, num)
	}

	maxSub := length / 4

	rem := length % 4
	switch rem {
		case 0:
			l1 = maxSub
			l2 = maxSub * 2
			l3 = maxSub * 3
		case 1:
			l1 = maxSub + 1
			l2 = maxSub*2 + 1
			l3 = maxSub*3 + 1
		case 2:
			l1 = maxSub + 1
			l2 = maxSub*2 + 2
			l3 = maxSub*3 + 2
		case 3:
			l1 = maxSub + 1
			l2 = maxSub*2 + 2
			l3 = maxSub*3 + 3
	}

	slc1 = integerList[:l1]
	slc2 = integerList[l1:l2]
	slc3 = integerList[l2:l3]
	slc4 = integerList[l3:]

	c := make(chan []int)
	go sortArr(slc1, c)
	go sortArr(slc2, c)
	go sortArr(slc3, c)
	go sortArr(slc4, c)

	arr1:= <- c
	arr2:= <- c
	arr3:= <- c
	arr4:= <- c

	sortedArr := make([]int, length)
	
	for i, v := range arr1 {
		sortedArr[i] = v
	}
	
	for i, v := range arr2 {
		sortedArr[i+len(arr1)] = v
	}
	
	for i, v := range arr3 {
		sortedArr[i+len(arr1)+len(arr2)] = v
	}
	
	for i, v := range arr4 {
		sortedArr[i+len(arr1)+len(arr2)+len(arr3)] = v
	}

	sort.Ints(sortedArr)
	fmt.Println("Final sorted list: ", sortedArr)
}