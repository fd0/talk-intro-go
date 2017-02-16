package main

import "fmt"

// START OMIT
func main() {
	var a [4]int

	a = [4]int{0, 1, 1, 2}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a[2])

	var b = [...]int{0, 1, 1, 2, 3, 5, 8}
	fmt.Printf("%#v\n", b)
	fmt.Println(len(b))
}

// END OMIT
