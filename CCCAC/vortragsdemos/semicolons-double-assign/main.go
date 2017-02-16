package main

import "fmt"

func main() {
	var x, y int

	x = 23; y = 42
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
