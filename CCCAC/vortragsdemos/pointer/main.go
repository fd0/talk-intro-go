package main

import "fmt"

// START OMIT
func double(a int) int {
	return 2 * a
}

func increment(a *int) {
	*a += 1
}

func main() {
	var x = 23

	fmt.Println(double(x))

	fmt.Println(x)

	increment(&x)
	fmt.Println(x)
}
// END OMIT
