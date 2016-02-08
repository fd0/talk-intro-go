package main

import "fmt"

func init() {
	fmt.Println("This is init1")
}

func init() {
	fmt.Println("This is init2")
}

func main() {
	fmt.Println("This is main")
}
