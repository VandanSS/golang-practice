package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)


type ChopS struct {
}


type Philosopher struct {
	id int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(ChopS)
	},
}

var w sync.WaitGroup

func (p Philosopher) eat(host chan int) {

	n := rand.Intn(8)
	defer w.Done()
	
	<-host
	fmt.Printf("starting to eat %d\n", p.id)

	left := pool.Get()
	right := pool.Get()
	
	pool.Put(left)
	pool.Put(right)
	
	fmt.Printf("finishing eating %d\n", p.id)
	host <- 1
	time.Sleep(time.Duration(n)*time.Second)
}

func main() {

	// Make a host
	host := make(chan int, 2)

	// Put 5 chopsticks in pool
	for i := 0; i < 5; i++ {
		pool.Put(new(ChopS))
	}

	// Make 5 philosophers
	philosophers := make([]*Philosopher, 5)
	
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1}
	}


	for i := 0; i < 5; i++ {
		// Each philosophers should eat 3 times
		for j := 0; j < 3; j++ {
			w.Add(1)
			go philosophers[i].eat(host)
		}	
	}

	host <- 1
	host <- 1
	w.Wait()
}