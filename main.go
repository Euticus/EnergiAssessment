package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}

	}
}

var nthPrime int = 1000000

func main() {

	var prime int
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i <= nthPrime; i++ {
		prime = <-ch
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Printf("Last element: %v\n", prime)
}
