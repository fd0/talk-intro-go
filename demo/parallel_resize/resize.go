package main

import (
	"fmt"
	"os"
	"os/exec"
)

func resize(x, y int, input, output string) error {
	res := fmt.Sprintf("%dx%d", x, y)
	cmd := exec.Command("convert", input, "-resize", res, output)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
