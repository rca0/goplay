// fetch shows the content finded in each custom URL

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\v", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: while reading %s: %v\n", url, err)
		}

		fmt.Printf("%s", b)
	}
}
