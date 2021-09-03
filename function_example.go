package main

import (
	"fmt"
	"time"
)

func DisplayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}
