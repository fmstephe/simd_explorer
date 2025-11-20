package vinsertf128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vinsertf128_256_one.s
var assemblyVinsertf128256One string

//go:embed stub_vinsertf128_256_one.go
var stubVinsertf128256One string

type VINSERTF128256ONE struct {
}

func (v *VINSERTF128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VINSERTF128256ONE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VINSERTF128256ONE) Name() string {
	return "VINSERTF128 (256 bit) one"
}

func (v *VINSERTF128256ONE) Description() string {
	return "Insert 128-bit block into upper lane (imm8=1): dst[255:128]=block, dst[127:0]=base[127:0]."
}

func (v *VINSERTF128256ONE) Stub() string {
	return stubVinsertf128256One
}

func (v *VINSERTF128256ONE) Assembly() string {
	return assemblyVinsertf128256One
}

func (v *VINSERTF128256ONE) Run(inputs [][]byte) (output []byte) {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(inputs[0]))
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vinsertf128256One(&base, &block, &ret)

	log.Printf("VINSERTF128256ONE base %v block %v output %v", base, block, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VINSERTF128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
