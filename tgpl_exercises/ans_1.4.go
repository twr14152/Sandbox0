package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//new nested map
	filenames := make(map[string]map[string]int)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//added filenames in countlines
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	fmt.Println("********************** RESULTS ************************")
	for line, n := range counts {
		if n > 1 {
			fmt.Println()
			fmt.Printf("\t%d\t%s", n, line)
			for files, fc := range filenames[line] {
				fmt.Printf("\t%d \t\t\t\t%s", fc, files)
			}
			fmt.Println()
		}
	}
}
func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fn := f.Name()
		counts[text]++
		if filenames[text] == nil {
			filenames[text] = make(map[string]int)
		}
		filenames[text][fn]++
	}
