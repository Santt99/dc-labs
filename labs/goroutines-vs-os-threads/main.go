package main

import "fmt"

func routines(w chan int) {
	var quantity int
	quantity = <-w
	w <- (quantity + 1)
	fmt.Printf("Goroutine: %d\n", quantity+1)
	go routines(w)
	for {
	}
}
func main() {
	var counter chan int
	counter = make(chan int)
	// counter <- 1
	routines(counter)
	close(counter)
}
