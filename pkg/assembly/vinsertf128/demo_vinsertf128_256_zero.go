package vinsertf128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vinsertf128_256_zero.s
var assemblyVinsertf128256Zero string

//go:embed stub_vinsertf128_256_zero.go
var stubVinsertf128256Zero string

type VINSERTF128256ZERO struct {
}

func (v *VINSERTF128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VINSERTF128256ZERO) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VINSERTF128256ZERO) Name() string {
	return "VINSERTF128 (256 bit) zero"
}

func (v *VINSERTF128256ZERO) Description() string {
	return "Insert 128-bit block into lower lane (imm8=0): dst[127:0]=block, dst[255:128]=base[255:128]."
}

func (v *VINSERTF128256ZERO) Stub() string {
	return stubVinsertf128256Zero
}

func (v *VINSERTF128256ZERO) Assembly() string {
	return assemblyVinsertf128256Zero
}

func (v *VINSERTF128256ZERO) Run(inputs [][]byte) (output []byte) {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(inputs[0]))
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vinsertf128256Zero(&base, &block, &ret)

	log.Printf("VINSERTF128256ZERO base %v block %v output %v", base, block, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VINSERTF128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
