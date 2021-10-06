package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balancce int64

func crediit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// adds 100 to balance atomically
		atomic.AddInt64(&balancce, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func debiit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		atomic.AddInt64(&balancce, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func synch() {
	var wg sync.WaitGroup
	balancce = 200
	fmt.Println("Initial balance is", balancce)
	wg.Add(1) // add 1 to the WaitGroup counter
	go crediit(&wg)
	wg.Add(1) // add 1 to the WaitGroup counter
	go debiit(&wg)
	wg.Wait() // blocks until WaitGroup counter is 0
	fmt.Println("Final balance is", balancce)
}
