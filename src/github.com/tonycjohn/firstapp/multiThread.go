package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func multiThread() {
	//runtime.GOMAXPROCS(100) //sets max number of threads to 100
	fmt.Printf("Threads:%v\n", runtime.GOMAXPROCS(-1))
	fmt.Printf("Launching multiple Threads\n")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go goRoutineTest(i)
		go increment()
	}
	wg.Wait()
	fmt.Printf("Launched multiple Threads\n")
}

func goRoutineTest(val int) {
	fmt.Printf("Thread no:%v\n", val)
	wg.Done()
}

func increment() {
	counter++
	fmt.Printf("counter value %v\n", counter)
	wg.Done()
}
