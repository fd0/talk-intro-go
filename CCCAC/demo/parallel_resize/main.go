package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func errx(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: inputdir outputdir\n")
		os.Exit(1)
	}

	input, output := os.Args[1], os.Args[2]
	fmt.Printf("input %v, output %v\n", input, output)

	inputChannel := make(chan string)
	outputChannel := make(chan string)

	go func() {
		dir, err := os.Open(input)
		errx(err)

		files, err := dir.Readdirnames(-1)
		errx(err)

		for _, file := range files {
			inputChannel <- filepath.Join(input, file)
		}
		close(inputChannel)
	}()

	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go resizeWorker(&wg, output, inputChannel, outputChannel)
	}

	go func() {
		for file := range outputChannel {
			fmt.Printf("done: %v\n", file)
		}
	}()

	wg.Wait()
	close(outputChannel)
}
