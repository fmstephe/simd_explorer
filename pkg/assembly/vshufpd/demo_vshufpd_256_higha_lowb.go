package vshufpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufpd_256_higha_lowb.s
var assemblyVshufpd256Higha_lowb string

//go:embed stub_vshufpd_256_higha_lowb.go
var stubVshufpd256Higha_lowb string

type VSHUFPD256HIGHA_LOWB struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPD256HIGHA_LOWB() *VSHUFPD256HIGHA_LOWB {
	return &VSHUFPD256HIGHA_LOWB{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VSHUFPD256HIGHA_LOWB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPD256HIGHA_LOWB) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPD256HIGHA_LOWB) Name() string {
	return "VSHUFPD (256 bit) higha_lowb"
}

func (v *VSHUFPD256HIGHA_LOWB) Description() string {
	return "Shuffle pairs per 128-bit lane: ret = [vals1[1], vals2[0], vals1[3], vals2[2]] (imm8=0x01)."
}

func (v *VSHUFPD256HIGHA_LOWB) Stub() string {
	return stubVshufpd256Higha_lowb
}

func (v *VSHUFPD256HIGHA_LOWB) Assembly() string {
	return assemblyVshufpd256Higha_lowb
}

func (v *VSHUFPD256HIGHA_LOWB) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vshufpd256Higha_lowb(&vals1, &vals2, &ret)

	log.Printf("VSHUFPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VSHUFPD256HIGHA_LOWB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
