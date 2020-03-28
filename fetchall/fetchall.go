// find URL concorrency and show waste time and size

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // create channel

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // init goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from the channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close() // avoid leakage of resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
