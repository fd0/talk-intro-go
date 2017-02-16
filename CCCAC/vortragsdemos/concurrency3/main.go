package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START worker OMIT
func print(worker, count int, ch chan time.Duration, done <-chan struct{}) {
	start := time.Now()
	defer func() {
		ch <- time.Since(start)
	}()

	for count > 0 {
		fmt.Printf("[worker %3d] %3d bottles of mate on the wall\n", worker, count)
		timeout := time.After(time.Duration(count) * 50 * time.Millisecond)
		count--
		select {
		case <-done:
			return
		case <- timeout:
		}
	}
}

// END worker OMIT

// START main OMIT
func main() {
	ch := make(chan time.Duration)
	done := make(chan struct{})

	for i := 0; i < 15; i++ {
		go print(i, rand.Intn(20), ch, done)
	}

	go func() {
		<-time.After(5 * time.Second)
		close(done)
	}()

	for i := 0; i < 15; i++ {
		d := <-ch
		fmt.Printf("print() finished after %v\n", d)
	}
}

// END main OMIT
