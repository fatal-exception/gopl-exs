// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
	"strings"
)

func getUrls(filename string) []string {
	var urls []string
	f, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			url := strings.Split(line, ",")[1]
			urls = append(urls, "http://" + url)
		}
	}
	return urls
}

func fain2() {
	start := time.Now()
	filename := os.Args[1]
	urls := getUrls(filename)
	ch := make(chan string)
	for _, url := range urls {
		//fmt.Println("Starting a fetch of " + url)
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
