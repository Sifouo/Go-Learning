package main

import (
	"fmt"
	"time"
)

func DisplayTime() {
	fmt.Println(time.Now())
	for i := 4; i >= 0; i-- {
		fmt.Println(i)
	}
Outerloop: //outer is a label for two for loops
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break Outerloop //outer is a label
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----------")
	}
}
