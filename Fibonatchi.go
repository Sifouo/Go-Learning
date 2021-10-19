package main

import (
	"fmt"
	"time"
)

type index struct {
	i   int
	fib int
}

func fib(n int, c chan index) {
	a, b := 1, 1
	for i := 0; i <= n; i++ {

		c <- index{i: i, fib: a}
		a, b = b, a+b
		time.Sleep(1 * time.Second)

	}
	close(c)
}

func main() {
	c := make(chan index)
	var n int
	fmt.Println("inter the number of index:   ")
	fmt.Scanln(&n)
	go fib(n, c)

	for i := range c {
		fmt.Println(i.i, ":", i.fib)
	}
}
