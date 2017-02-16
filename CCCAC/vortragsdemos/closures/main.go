package main

import "fmt"

// START OMIT
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	gen := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(gen(), gen(), gen(), gen())
	}
}
// END OMIT
