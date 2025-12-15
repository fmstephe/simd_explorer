package vshufpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufpd_256_lowa_lowb.s
var assemblyVshufpd256Lowa_lowb string

//go:embed stub_vshufpd_256_lowa_lowb.go
var stubVshufpd256Lowa_lowb string

type VSHUFPD256LOWA_LOWB struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPD256LOWA_LOWB() *VSHUFPD256LOWA_LOWB {
	return &VSHUFPD256LOWA_LOWB{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VSHUFPD256LOWA_LOWB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPD256LOWA_LOWB) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPD256LOWA_LOWB) Name() string {
	return "VSHUFPD (256 bit) lowa_lowb"
}

func (v *VSHUFPD256LOWA_LOWB) Description() string {
	return "Shuffle pairs per 128-bit lane: ret = [vals1[0], vals2[0], vals1[2], vals2[2]] (imm8=0x00)."
}

func (v *VSHUFPD256LOWA_LOWB) Stub() string {
	return stubVshufpd256Lowa_lowb
}

func (v *VSHUFPD256LOWA_LOWB) Assembly() string {
	return assemblyVshufpd256Lowa_lowb
}

func (v *VSHUFPD256LOWA_LOWB) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vshufpd256Lowa_lowb(&vals1, &vals2, &ret)

	log.Printf("VSHUFPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VSHUFPD256LOWA_LOWB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
