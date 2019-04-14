package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 uses a loop to do a PopCount. Ex2.3
func PopCount2(x uint64) int {
	var result int
	for i := range pc {
		result += int(pc[byte(x>>(uint(i)*8))])
	}
	return result

}

// PopCount3 shifts each individual bit and counts it, for dreadful performance
func PopCount3(x uint64) int {
	var result int
	var i uint
	for i = 0; i < 64; i++ {
		if x&1 == 1 {
			result++
		}
		x = x >> 1
	}
	return result
}

// PopCount4 uses negation to count a population. Ex2.5
func PopCount4(x uint64) int {
	var result int
	afterX := x & (x - 1)
	for afterX != x {
		result++
		x = afterX
		afterX = x & (x - 1)
	}
	return result
}
