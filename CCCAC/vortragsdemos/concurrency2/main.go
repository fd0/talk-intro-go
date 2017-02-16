package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func print(worker, count int, ch chan time.Duration) {
	start := time.Now()
	for count > 0 {
		fmt.Printf("[worker %3d] %3d bottles of mate on the wall\n", worker, count)
		time.Sleep(time.Duration(count)*time.Millisecond)
		count--
	}

	ch <- time.Since(start)
}

func main() {
	ch := make(chan time.Duration)
	for i := 0; i < 15; i++ {
		go print(i, rand.Intn(20), ch)
	}

	for i := 0; i < 15; i++ {
		d := <-ch
		fmt.Printf("print() finished after %v\n", d)
	}
}

// END OMIT
