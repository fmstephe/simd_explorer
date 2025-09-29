package register

import "fmt"

func partsForPartSize(partSize, totalSize int) int {
	mustValidPartSize(partSize)
	mustValidInputOutputSize(totalSize)
	return totalSize / partSize
}

func mustValidInputOutputSize(totalSize int) {
	switch totalSize {
	case 8, 16, 32, 64, 128, 256, 512:
	default:
		panic(fmt.Errorf("Unsupported input/output size: %d", totalSize))
	}
}

func mustValidPartSize(partSize int) {
	switch partSize {
	case 8, 16, 32, 64:
	default:
		panic(fmt.Errorf("Unsupported bitsize value: %d", partSize))
	}
}

func mustValidSimdsize(simdSize int) {
	switch simdSize {
	case 128, 256, 512:
	default:
		panic(fmt.Errorf("Unsupported simdSize value: %d", simdSize))
	}
}
