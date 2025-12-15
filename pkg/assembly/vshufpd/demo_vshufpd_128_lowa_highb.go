package vshufpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufpd_128_lowa_highb.s
var assemblyVshufpd128Lowa_highb string

//go:embed stub_vshufpd_128_lowa_highb.go
var stubVshufpd128Lowa_highb string

type VSHUFPD128LOWA_HIGHB struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPD128LOWA_HIGHB() *VSHUFPD128LOWA_HIGHB {
	return &VSHUFPD128LOWA_HIGHB{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VSHUFPD128LOWA_HIGHB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPD128LOWA_HIGHB) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPD128LOWA_HIGHB) Name() string {
	return "VSHUFPD (128 bit) lowa_highb"
}

func (v *VSHUFPD128LOWA_HIGHB) Description() string {
	return "Shuffle pairs: ret = [vals1[0], vals2[1]] (imm8=0x02)."
}

func (v *VSHUFPD128LOWA_HIGHB) Stub() string {
	return stubVshufpd128Lowa_highb
}

func (v *VSHUFPD128LOWA_HIGHB) Assembly() string {
	return assemblyVshufpd128Lowa_highb
}

func (v *VSHUFPD128LOWA_HIGHB) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vshufpd128Lowa_highb(&vals1, &vals2, &ret)

	log.Printf("VSHUFPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VSHUFPD128LOWA_HIGHB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
