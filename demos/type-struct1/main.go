package main

import "fmt"

type Sex int

type Gopher struct {
	name string
	sex  Sex
}

const (
	Unknown Sex = 0
	Female  Sex = 1
	Male    Sex = 2
)

func main() {
	g := Gopher{name: "Felix"}
	fmt.Println(g.name, g.sex)
}
