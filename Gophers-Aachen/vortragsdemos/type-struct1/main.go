package main

import "fmt"

// START OMIT
type Color int

const (
	Unknown Color = 0
	Blue    Color = 1
	Pink    Color = 2
	Purple  Color = 3
)

type Gopher struct {
	name  string
	color Color
}

func greet(g Gopher) {
	fmt.Println("Hello", g.name, "(", g.color, ")")
}

func main() {
	g := Gopher{name: "Felix"}
	fmt.Printf("%#v\n", g)
	g.color = Purple
	fmt.Printf("%#v\n", g)
	greet(g)
}

// END OMIT
