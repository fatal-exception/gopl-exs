package popcount

import (
	"fmt"
	"testing"
)

var implementations = map[string]func(uint64) int{
	"PopCount":  PopCount,
	"PopCount2": PopCount2,
	"PopCount3": PopCount3,
	"PopCount4": PopCount4,
}

func TestPopcount(t *testing.T) {
	testPopcount := func(t *testing.T, input uint64, want int, popcountImpl func(uint64) int) {
		t.Helper()
		retValue := popcountImpl(input)
		if retValue != want {
			t.Fatalf("%d is not %d", retValue, want)
		}
	}

	for name, function := range implementations {
		t.Run(fmt.Sprintf("ensure %s returns correct result", name), func(t *testing.T) {
			testPopcount(t, uint64(2), 1, function)
		})
	}
}

func BenchmarkPopcount(b *testing.B) {
	benchmarkPopcount := func(b *testing.B, popcountImpl func(uint64) int) {
		tester1 := uint64(367)
		for i := 0; i < b.N; i++ {
			popcountImpl(tester1)
		}
	}
	for name, function := range implementations {
		b.Run(fmt.Sprintf("benchmark %s", name), func(b *testing.B) {
			benchmarkPopcount(b, function)
		})
	}
}
