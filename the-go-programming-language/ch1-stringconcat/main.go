package main

import (
	"fmt"
	"strings"
	"time"
)

func slow(xs []string) string {
	s, sep := "", ""
	for _, arg := range xs {
		s += sep + arg
		sep = " "
	}
	return s
}

func fast(xs []string) string {
	return strings.Join(xs, " ")
}

func buildArr(N int) []string {
	arr := make([]string, N)
	for i := range arr {
		arr[i] = "asdf"
	}
	return arr
}

func main() {

	var t0 time.Time
	var elapsed time.Duration
	for N := 100; N < 1000000; N *= 10 {
		arr := buildArr(N)

		t0 = time.Now()
		fast(arr)
		elapsed = time.Since(t0)
		fmt.Println("Fast time: ", elapsed)

		t0 = time.Now()
		slow(arr)
		elapsed = time.Since(t0)
		fmt.Println("Slow time: ", elapsed)
	}
}
