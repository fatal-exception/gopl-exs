package main

import (
	"github.com/fatal-exception/gopl-exs/ch2/popcount"
	"fmt"
)

func main() {
	var i uint64 = 7
	fmt.Println(popcount.PopCount(i))
	fmt.Println(popcount.PopCount2(i))
}
