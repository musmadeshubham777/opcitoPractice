package main

import (
	"fmt"
	"sync"
)

//Critical Section : When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time.

// mutex is used to  avoid race condition
// A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.
// Mutex is available in the sync package. There are two methods defined on Mutex namely
// 1.Lock 2.Unlock.

var x = 0

func increment(wg *sync.WaitGroup, c chan bool) {
	c <- true
	x = x + 1
	<-c
	wg.Done()
}

func mainMutex() {

	var wg sync.WaitGroup

	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}

	wg.Wait()
	fmt.Println("Final Value of x", x)
}
