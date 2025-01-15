package main

import (
	"fmt"
	"sync"
)

func main() {

}

func MutexDeadlock() {
	var wg sync.WaitGroup
	var mu1, mu2 sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("goroutine 1 acquired lock 1")
		mu2.Lock()
		fmt.Println("goroutine 1 acquired lock 2")
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("goroutine 2 acquired lock 2")
		mu1.Lock()
		fmt.Println("goroutine 2 acquired lock 1")
	}()
	wg.Wait()
}

func CircularWaitingBetweenGoroutines() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1 // Goroutine 1 waits for ch1 to be received
		<-ch2    // Goroutine 1 waits to receive from ch2
	}()

	go func() {
		ch2 <- 2 // Goroutine 2 waits for ch2 to be received
		<-ch1    // Goroutine 2 waits to receive from ch1
	}()

	// Both goroutines are waiting on each other, causing a deadlock.
	select {}
}

func MainGoroutineBlockingAllOthers() {
	ch := make(chan int)

	go func() {
		<-ch // Worker goroutine waiting for data from ch
	}()

	<-ch // Main goroutine also waiting for data from ch
	// Both goroutines are blocked, leading to a deadlock.
}

func MisuseOfWaitGroupWithChannels() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		ch <- 42 // This goroutine is waiting for a receiver
	}()

	wg.Wait() // Main goroutine waiting for worker to finish
	<-ch      // But this line is never reached, causing a deadlock
}

func NoReceiversForUnbufferedChannels() {
	ch := make(chan int)

	go func() {
		ch <- 42 // This goroutine blocks because there is no receiver
	}()

	ch <- 1 // The main goroutine also blocks here, leading to deadlock
}
