package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastsd_256_m64.s
var assemblyVbroadcastsd256M64 string

//go:embed stub_vbroadcastsd_256_m64.go
var stubVbroadcastsd256M64 string

type VBROADCASTSD256M64 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVBROADCASTSD256M64() *VBROADCASTSD256M64 {
	return &VBROADCASTSD256M64{
		scalar: number.NewNamedFloatParameter("scalar", 64, 64),
		ret:    number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBROADCASTSD256M64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
	}
}

func (v *VBROADCASTSD256M64) Output() *number.Parameter {
	return v.ret
}

func (v *VBROADCASTSD256M64) Name() string {
	return "VBROADCASTSD (256 bit) m64"
}

func (v *VBROADCASTSD256M64) Description() string {
	return "Broadcast 64-bit scalar from memory into all 4 lanes of YMM."
}

func (v *VBROADCASTSD256M64) Stub() string {
	return stubVbroadcastsd256M64
}

func (v *VBROADCASTSD256M64) Assembly() string {
	return assemblyVbroadcastsd256M64
}

func (v *VBROADCASTSD256M64) Run(_ [][]byte) (output []byte) {
	scalar := number.ToFloat64(v.scalar.FlatData())

	ret := [4]float64{}

	vbroadcastsd256M64(&scalar, &ret)

	log.Printf("VBROADCASTSD256M64 scalar %v output %v", scalar, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VBROADCASTSD256M64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
