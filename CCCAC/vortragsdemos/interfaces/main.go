package main

import "fmt"

type Color int

const (
	Unknown Color = 0
	Blue    Color = 1
	Pink    Color = 2
	Purple  Color = 3
)

// START Color OMIT
func (c Color) String() string {
	switch c {
	case Unknown:
		return "unknown"
	case Blue:
		return "blue"
	case Pink:
		return "pink"
	case Purple:
		return "purple"
	default:
		return "other"
	}
}

// END Color OMIT

type Gopher struct {
	name  string
	color Color
}

// START main OMIT
func (g Gopher) greet() {
	fmt.Println("Hello", g.name, "(", g.color, ")")
}

func main() {
	g := Gopher{name: "Larissa", color: Blue}
	g.greet()
}

// END main OMIT
