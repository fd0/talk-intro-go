package main

import "fmt"

// START OMIT
const mask = (1 << 128) - 1

func main() {
	var m1 float64 = mask
	var m2 uint32 = mask & 0xffffffff

	fmt.Println(m1, m2)
}
// END OMIT
