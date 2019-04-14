package popcount

import (
	"fmt"
	"testing"
)

func TestPopcount(t *testing.T) {
	testPopcount := func(t *testing.T, input uint64, want int, popcountImpl func(uint64) int) {
		t.Helper()
		retValue := popcountImpl(input)
		if retValue != want {
			t.Fatalf("%d is not %d", retValue, want)
		}
	}

	implementations := map[string]func(uint64) int{
		"PopCount":  PopCount,
		"PopCount2": PopCount2,
		"PopCount3": PopCount3,
	}

	for name, function := range implementations {
		t.Run(fmt.Sprintf("ensure %s returns correct result", name), func(t *testing.T) {
			testPopcount(t, uint64(2), 1, function)
		})
	}
}

func BenchmarkPopcount(b *testing.B) {
	b.Run("popcount 1", func(b *testing.B) {
		tester1 := uint64(367)
		for i := 0; i < b.N; i++ {
			PopCount(tester1)
		}
	})
	b.Run("popcount 2", func(b *testing.B) {
		tester1 := uint64(367)
		for i := 0; i < b.N; i++ {
			PopCount2(tester1)
		}
	})
	b.Run("popcount 3", func(b *testing.B) {
		tester1 := uint64(367)
		for i := 0; i < b.N; i++ {
			PopCount3(tester1)
		}
	})
}
