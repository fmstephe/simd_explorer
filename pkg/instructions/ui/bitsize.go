package ui

import "fmt"

func textWidthForBitsize(bitsize int) int {
	mustValidBitsize(bitsize)
	switch bitsize {
	case 8:
		return 4
	case 16:
		return 6
	case 32:
		return 11
	case 64:
		return 21
	}
	panic("Unreachable")
}

func partsForBitsize(bitsize, simdsize int) int {
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
