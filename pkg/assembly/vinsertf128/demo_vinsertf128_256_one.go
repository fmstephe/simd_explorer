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
	vals  *number.Parameter
	block *number.Parameter
	ret   *number.Parameter
}

func NewVINSERTF128256ONE() *VINSERTF128256ONE {
	return &VINSERTF128256ONE{
		vals:  number.NewNamedFloatParameter("vals", 256, 32),
		block: number.NewNamedFloatParameter("block", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VINSERTF128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.block,
	}
}

func (v *VINSERTF128256ONE) Output() *number.Parameter {
	return v.ret
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

func (v *VINSERTF128256ONE) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(v.block.FlatData()))

	ret := [8]float32{}

	vinsertf128256One(&vals, &block, &ret)

	log.Printf("VINSERTF128256ONE vals %v block %v output %v", vals, block, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VINSERTF128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
