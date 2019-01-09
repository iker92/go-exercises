package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) > 0 {
		for _, fileIn := range files {
			file, err := os.Open(fileIn)
			if err != nil {
				fmt.Printf("Error during opening file : %v", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	for file := range counts {
		for line, n := range counts[file] {
			if n > 1 {
				fmt.Printf("Word %v occured %d times in file %v \n", line, n, file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	counts[f.Name()] = make(map[string]int)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
