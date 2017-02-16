package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func print(wg *sync.WaitGroup, worker, count int) {
	defer wg.Done()
	fmt.Printf("[worker %3d] %3d bottles of mate on the wall\n", worker, count)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go print(&wg, i, rand.Intn(20))
	}

	wg.Wait()
}

// END OMIT
