package main

import "fmt"

func main() {
	var x int = 23
	if x == 17 {
		fmt.Println("x is 17")
	} else {
		fmt.Println("x is not 17")
	}

	var i int
	for i = 1; i < 40; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FIZZBUZZ")
		case i%3 == 0:
			fmt.Println("FIZZ")
		case i%5 == 0:
			fmt.Println("BUZZ")
		default:
			fmt.Println(i)
		}
	}
}
