package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"golang.org/x/sys/cpu"
)

//go:embed asm_128_k.s
var assembly128K string

//go:embed stub_128_k.go
var stub128K string

type Vpbroadcastb128K struct {
}

func (v *Vpbroadcastb128K) InputSizes() []int {
	return []int{8, 64}
}

func (v *Vpbroadcastb128K) OutputSize() int {
	return 128
}

func (v *Vpbroadcastb128K) Name() string {
	return "VPBROADCASTB XMM (128K bit) K"
}

func (v *Vpbroadcastb128K) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb128K) Stub() string {
	return stub128K
}

func (v *Vpbroadcastb128K) Assembly() string {
	return assembly128K
}

func (v *Vpbroadcastb128K) Run(inputs [][]byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128K(inputs[0][0], asmutil.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *Vpbroadcastb128K) Supported() bool {
	// Requires: AVX, AVX2, AVX512F, AVX512VL, SSE2
	return cpu.X86.HasAVX &&
		cpu.X86.HasAVX2 &&
		cpu.X86.HasAVX512F &&
		cpu.X86.HasAVX512VL &&
		cpu.X86.HasSSE2
}
