package main

import "fmt"

// START OMIT
func main() {
	var s []int
	fmt.Printf("%v\n", s)

	s = make([]int, 3)
	s[0] = 5
	fmt.Printf("%v\n", s)

	s = append(s, 23)
	fmt.Printf("%v\n", s)
}
// END OMIT
