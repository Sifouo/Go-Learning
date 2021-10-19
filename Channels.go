package main

import (
	"fmt"
	"time"
)

func channels() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	fmt.Scanln()

}

//////////////////////////////////
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	/////////////////////////
	//time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel ...")
}

///////////////////////////////////
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from channel:", <-ch)

}
