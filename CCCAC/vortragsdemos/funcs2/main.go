package main

import (
	"errors"
	"fmt"
	"math"
)

// START OMIT
func split(c complex128) (r, i float64) {
	return real(c), imag(c)
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, errors.New("Negative number")
	}

	return math.Sqrt(i), nil
}

func main() {
	c := complex(23, 42)
	fmt.Println(split(c))

	fmt.Println(sqrt(25))
	s, err := sqrt(-25)
	fmt.Println(s, err)
}
// END OMIT
