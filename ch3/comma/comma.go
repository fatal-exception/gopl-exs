package comma

import (
	"bytes"
	"strconv"
)

// Comma inserts commas into a number
// Ex 3.10
func Comma(num int) string {
	var buf bytes.Buffer
	numString := strconv.Itoa(num)
	for idx, letter := range numString {
		buf.WriteRune(letter)
		if (idx != len(numString)-1) && ((len(numString)-idx)%3 == 1) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
