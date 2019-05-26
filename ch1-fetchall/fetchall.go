package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	stdoutCh := make(chan string)
	fileCh := make(chan []byte)
	outFile, err := os.Create("fetchlog.txt")
	if err != nil {
		fmt.Printf("opening log file: %s\n", err)
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		go fetch(url, stdoutCh, fileCh)
	}
	for range os.Args[1:] {
		_, err := outFile.Write(<-fileCh)
		if err != nil {
			fmt.Printf("writing to log file: %s\n", err)
		}
		fmt.Println(<-stdoutCh)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, stdoutCh chan<- string, fileCh chan<- []byte) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fileCh <- []byte(fmt.Sprint(err))
		stdoutCh <- fmt.Sprint(err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fileCh <- []byte(fmt.Sprintf("while reading %s: %v", url, err))
		stdoutCh <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fileCh <- buf
	stdoutCh <- fmt.Sprintf("%.2fs  %7d  %s", secs, len(buf), url)
}
