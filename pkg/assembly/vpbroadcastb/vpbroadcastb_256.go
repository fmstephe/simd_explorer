package vpbroadcastb

import (
	_ "embed"

	"golang.org/x/sys/cpu"
)

//go:embed asm_256.s
var assembly256 string

//go:embed stub_256.go
var stub256 string

type Vpbroadcastb256 struct {
}

func (v *Vpbroadcastb256) InputSizes() []int {
	return []int{8}
}

func (v *Vpbroadcastb256) OutputSize() int {
	return 256
}

func (v *Vpbroadcastb256) Name() string {
	return "VPBROADCASTB YMM (256 bit)"
}

func (v *Vpbroadcastb256) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb256) Stub() string {
	return stub256
}

func (v *Vpbroadcastb256) Assembly() string {
	return assembly256
}

func (v *Vpbroadcastb256) Run(inputs [][]byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb256(inputs[0][0], &ret)
	return ret[:]
}

func (v *Vpbroadcastb256) Supported() bool {
	// Requires: AVX, AVX2, SSE2
	return cpu.X86.HasAVX &&
		cpu.X86.HasAVX2 &&
		cpu.X86.HasSSE2
}
