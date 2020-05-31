package main

import (
	"fmt"
	"time"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 0
	c2 <- 1
	return func() int {
		go sum(c1, c2)
		time.Sleep(time.Second)
		return <-c1
	}
}

func sum(c1 chan int, c2 chan int) {
	x := <-c1
	y := <-c2
	c1 <- y
	c2 <- x + y
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
