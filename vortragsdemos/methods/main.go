package main

import "fmt"

type Color int

const (
	Unknown Color = 0
	Blue    Color = 1
	Pink    Color = 2
	Purple  Color = 3
)

// START OMIT
type Gopher struct {
	name  string
	color Color
}

func (g Gopher) greet() {
	fmt.Println("Hello", g.name, "(", g.color, ")")
}

func main() {
	g := Gopher{name: "Larissa", color: Blue}
	g.greet()
}

// END OMIT
