package ui

import "fmt"

func textWidthForBitsize(bitsize int) int {
	mustValidBitsize(bitsize)
	switch bitsize {
	case 8:
		return 3
	case 16:
		return 5
	case 32:
		return 10
	case 64:
		return 20
	}
	panic("Unreachable")
}

func inputsForBitsize(bitsize, simdsize int) int {
	mustValidBitsize(bitsize)
	mustValidSimdsize(simdsize)
	return simdsize / bitsize
}

func mustValidBitsize(bitsize int) {
	switch bitsize {
	case 8, 16, 32, 64:
	default:
		panic(fmt.Errorf("Unsupported bitsize value: %d", bitsize))
	}
}

func mustValidSimdsize(simdsize int) {
	switch simdsize {
	case 128, 256, 512:
	default:
		panic(fmt.Errorf("Unsupported simdsize value: %d", simdsize))
	}
}
