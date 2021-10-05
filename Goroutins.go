package main

import (
	"fmt"
	"time"
)

func say(s string, t int) {
	for i := 0; i <= t; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func goruotin() {
	go say("hello", 2)
	go say("world", 3)
	fmt.Scanln()
}
