package main

import (
	"fmt"
	"path/filepath"
	"sync"
)

func resizeWorker(wg *sync.WaitGroup, outdir string, input <-chan string, output chan<- string) {
	defer wg.Done()

	for infile := range input {
		outfile := filepath.Join(outdir, filepath.Base(infile))
		err := resize(1024, 768, infile, outfile)
		if err != nil {
			fmt.Printf("error! %v\n", err)
		}
		output <- outfile
	}
}
