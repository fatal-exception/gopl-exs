// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("%v%v", "http://", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("\n%s\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "We have a problem reading %v: %v", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
