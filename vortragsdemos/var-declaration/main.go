package main

import (
	"fmt"
	"strings"
)

const salutation = "Dear Gopher"

func main() {
	var s = salutation + " Felix"
	fmt.Println(s)

	s2 := strings.ToLower(s)
	fmt.Println(s2)
}
