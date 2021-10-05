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
	say("hello", 3)
	say("world", 2)
}
