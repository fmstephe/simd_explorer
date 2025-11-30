package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastss_256_m32.s
var assemblyVbroadcastss256M32 string

//go:embed stub_vbroadcastss_256_m32.go
var stubVbroadcastss256M32 string

type VBROADCASTSS256M32 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVBROADCASTSS256M32() *VBROADCASTSS256M32 {
	return &VBROADCASTSS256M32{
		scalar: number.NewNamedFloatParameter("scalar", 32, 32),
		ret:    number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBROADCASTSS256M32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
	}
}

func (v *VBROADCASTSS256M32) Output() *number.Parameter {
	return v.ret
}

func (v *VBROADCASTSS256M32) Name() string {
	return "VBROADCASTSS (256 bit) m32"
}

func (v *VBROADCASTSS256M32) Description() string {
	return "Broadcast 32-bit scalar from memory into all 8 lanes of YMM."
}

func (v *VBROADCASTSS256M32) Stub() string {
	return stubVbroadcastss256M32
}

func (v *VBROADCASTSS256M32) Assembly() string {
	return assemblyVbroadcastss256M32
}

func (v *VBROADCASTSS256M32) Run() {
	scalar := number.ToFloat32(v.scalar.FlatData())

	ret := [8]float32{}

	vbroadcastss256M32(&scalar, &ret)

	log.Printf("VBROADCASTSS256M32 scalar %v output %v", scalar, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VBROADCASTSS256M32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
