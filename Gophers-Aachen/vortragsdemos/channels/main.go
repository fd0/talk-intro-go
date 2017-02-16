package main

import "fmt"

type Color int

const (
	Unknown Color = 0
	Blue    Color = 1
	Pink    Color = 2
	Purple  Color = 3
)

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

type Gopher struct {
	name  string
	color Color
}

func (g Gopher) greet() {
	fmt.Println("Hello", g.name, "(", g.color, ")")
}

// START OMIT
func main() {
	ch := make(chan Gopher)
	ch <- Gopher{name: "Emma", color: Pink}

	g := <-ch
	g.greet()
}

// END OMIT
