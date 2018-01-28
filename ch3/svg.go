// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"net/http"
	"log"
	"github.com/fatal-exception/gopl-exs/ch3/svg"
)

func main() {
	http.HandleFunc("/", svg.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

