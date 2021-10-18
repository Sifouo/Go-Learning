package main

func channels() {
	ch := make(chan int)
	ch <- 5
	value := <-ch

}
