// http is a parameterised lissajous server
package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// set defaults
	cycles := 5
	size := 100
	nframes := 64
	delay := 8

	for name, value := range map[string]*int{
		"cycles": &cycles,
		"size": &size,
		"nframes": &nframes,
		"delay": &delay} {
		if len(params[name]) > 0 {
			if num, err := strconv.Atoi(params[name][0]); err != nil {
				*value = num
			}
		}
	}

	lissajous(w, cycles, size, nframes, delay)
}

