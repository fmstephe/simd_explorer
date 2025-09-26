package vpbroadcastb

import (
	_ "embed"

	"golang.org/x/sys/cpu"
)

//go:embed asm_512.s
var assembly512 string

//go:embed stub_512.go
var stub512 string

type Vpbroadcastb512 struct {
}

func (v *Vpbroadcastb512) InputSizes() []int {
	return []int{8}
}

func (v *Vpbroadcastb512) OutputSize() int {
	return 512
}

func (v *Vpbroadcastb512) Name() string {
	return "VPBROADCASTB ZMM (512 bit)"
}

func (v *Vpbroadcastb512) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb512) Stub() string {
	return stub512
}

func (v *Vpbroadcastb512) Assembly() string {
	return assembly512
}

func (v *Vpbroadcastb512) Run(inputs [][]byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512(inputs[0][0], &ret)
	return ret[:]
}

func (v *Vpbroadcastb512) Supported() bool {
	// Requires: AVX, AVX2, AVX512F, AVX512VL, SSE2
	return cpu.X86.HasAVX &&
		cpu.X86.HasAVX2 &&
		cpu.X86.HasAVX512F &&
		cpu.X86.HasAVX512VL &&
		cpu.X86.HasSSE2
}
