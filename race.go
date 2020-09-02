package main

import (
	"fmt"
	"time"
)


var num int

func increment() {
	num++
	fmt.Println(num)
}


func decrement() {
	num--
	fmt.Println(num)
}


func main() {
	num = 200
	go increment()
	go decrement()

	time.Sleep(2*time.Second)
	fmt.Println("This program will exit as soon as main returns.")
}

/*

What is race condition?
Race condition is observed when multiple go routines are trying to access and modify the
same resource in this case the variable num in above code. This type of code results in 
varying output. The code outcome depends on interleavings so it defers all the time.

How can it occur?
It occurs due to communication between Goroutines through a shared variable.

In the above code there are two Goroutines 'increment' and 'decrement'. They increment
or decrement the value of variable num by 1. When this code is executed it prints two 
different outputs.

*/


/* Output

PS E:\go\src> go run .\race.go
200
201
This program will exit as soon as main returns.
PS E:\go\src> go run .\race.go
201
200
This program will exit as soon as main returns.
PS E:\go\src> go run .\race.go
200
201
This program will exit as soon as main returns.
PS E:\go\src>

*/