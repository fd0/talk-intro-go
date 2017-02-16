package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// START OMIT
	f, err := os.Open("Göthe-Faust.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	words := make(map[string]int)
	for sc.Scan() {
		if len(sc.Text()) < 6 { continue }
		words[sc.Text()] += 1
	}

	word, count := "", 0
	for w, c := range words {
		if c > count {
			word = w
			count = c
		}
	}

	fmt.Printf("most used word: %q (%d times)\n", word, words[word])
	// END OMIT
}
