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
func print(worker, count int) {
	fmt.Printf("[worker %3d] %3d bottles of mate on the wall\n", worker, count)
}

func main() {
	for i := 0; i < 15; i++ {
		print(i, rand.Intn(20))
	}
}

// END OMIT
