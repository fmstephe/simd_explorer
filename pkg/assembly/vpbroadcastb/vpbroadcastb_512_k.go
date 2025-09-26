package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"golang.org/x/sys/cpu"
)

//go:embed asm_512_k.s
var assembly512K string

//go:embed stub_512_k.go
var stub512K string

type Vpbroadcastb512K struct {
}

func (v *Vpbroadcastb512K) InputSizes() []int {
	return []int{8, 64}
}

func (v *Vpbroadcastb512K) OutputSize() int {
	return 512
}

func (v *Vpbroadcastb512K) Name() string {
	return "VPBROADCASTB ZMM (512K bit) K"
}

func (v *Vpbroadcastb512K) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb512K) Stub() string {
	return stub512K
}

func (v *Vpbroadcastb512K) Assembly() string {
	return assembly512K
}

func (v *Vpbroadcastb512K) Run(inputs [][]byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512K(inputs[0][0], asmutil.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *Vpbroadcastb512K) Supported() bool {
	// Requires: AVX, AVX2, AVX512F, AVX512VL, SSE2
	return cpu.X86.HasAVX &&
		cpu.X86.HasAVX2 &&
		cpu.X86.HasAVX512F &&
		cpu.X86.HasAVX512VL &&
		cpu.X86.HasSSE2
}
