package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastsf128_256_m128.s
var assemblyVbroadcastsf128256M128 string

//go:embed stub_vbroadcastsf128_256_m128.go
var stubVbroadcastsf128256M128 string

type VBROADCASTSF128256M128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVBROADCASTSF128256M128() *VBROADCASTSF128256M128 {
	return &VBROADCASTSF128256M128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBROADCASTSF128256M128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VBROADCASTSF128256M128) Output() *number.Parameter {
	return v.ret
}

func (v *VBROADCASTSF128256M128) Name() string {
	return "VBROADCASTSF128 (256 bit) m128"
}

func (v *VBROADCASTSF128256M128) Description() string {
	return "Broadcast 128-bit block (4x float32) from memory to both 128-bit lanes of YMM."
}

func (v *VBROADCASTSF128256M128) Stub() string {
	return stubVbroadcastsf128256M128
}

func (v *VBROADCASTSF128256M128) Assembly() string {
	return assemblyVbroadcastsf128256M128
}

func (v *VBROADCASTSF128256M128) Run() {
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vbroadcastsf128256M128(&block, &ret)

	log.Printf("VBROADCASTSF128256M128 block %v output %v", block, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VBROADCASTSF128256M128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
