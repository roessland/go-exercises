package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func filesToString(files map[string]bool) string {
	fs := []string{}
	for f := range files {
		fs = append(fs, f)
	}
	return strings.Join(fs, " ")
}

func main() {
	counts := make(map[string]int)
	var occurences map[string]map[string]bool
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurences, "")
	} else {
		occurences = make(map[string]map[string]bool)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurences, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			occurs := ""
			if occurences != nil {
				occurs = filesToString(occurences[line])
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, occurs)
		}
	}
}

func countLines(f *os.File, counts map[string]int, occurences map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if occurences != nil {
			if occurences[line] == nil {
				occurences[line] = make(map[string]bool)
			}
			occurences[line][filename] = true
		}
	}
}
