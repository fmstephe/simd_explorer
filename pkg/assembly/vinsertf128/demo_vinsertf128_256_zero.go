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
	vals  *number.Parameter
	block *number.Parameter
	ret   *number.Parameter
}

func NewVINSERTF128256ZERO() *VINSERTF128256ZERO {
	return &VINSERTF128256ZERO{
		vals:  number.NewNamedFloatParameter("vals", 256, 32),
		block: number.NewNamedFloatParameter("block", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VINSERTF128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.block,
	}
}

func (v *VINSERTF128256ZERO) Output() *number.Parameter {
	return v.ret
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

func (v *VINSERTF128256ZERO) Run(_ [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(v.block.FlatData()))

	ret := [8]float32{}

	vinsertf128256Zero(&vals, &block, &ret)

	log.Printf("VINSERTF128256ZERO vals %v block %v output %v", vals, block, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VINSERTF128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
