package main

import "fmt"

const salutation = "Dear friend" // HL
var name string

func main() {
	name = "Gopher" // HL

	fmt.Printf("%v %v\n", salutation, name)
}
