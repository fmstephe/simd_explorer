package vpbroadcastb

import (
	_ "embed"

	"golang.org/x/sys/cpu"
)

//go:embed asm_128.s
var assembly128 string

//go:embed stub_128.go
var stub128 string

type Vpbroadcastb128 struct {
}

func (v *Vpbroadcastb128) InputSize() int {
	return 8
}

func (v *Vpbroadcastb128) OutputSize() int {
	return 128
}

func (v *Vpbroadcastb128) Name() string {
	return "VPBROADCASTB XMM (128 bit)"
}

func (v *Vpbroadcastb128) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb128) Stub() string {
	return stub128
}

func (v *Vpbroadcastb128) Assembly() string {
	return assembly128
}

func (v *Vpbroadcastb128) Run(input []byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128(input[0], &ret)
	return ret[:]
}

func (v *Vpbroadcastb128) Supported() bool {
	// Requires: AVX, AVX2, SSE2
	return cpu.X86.HasAVX &&
		cpu.X86.HasAVX2 &&
		cpu.X86.HasSSE2
}
