package vunpcklpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklpd_128.s
var assemblyVunpcklpd128 string

//go:embed stub_vunpcklpd_128.go
var stubVunpcklpd128 string

type VUNPCKLPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKLPD128() *VUNPCKLPD128 {
	return &VUNPCKLPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VUNPCKLPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKLPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKLPD128) Name() string {
	return "VUNPCKLPD (128 bit) "
}

func (v *VUNPCKLPD128) Description() string {
	return "Unpack low doubles: ret = [vals1[0], vals2[0]]."
}

func (v *VUNPCKLPD128) Stub() string {
	return stubVunpcklpd128
}

func (v *VUNPCKLPD128) Assembly() string {
	return assemblyVunpcklpd128
}

func (v *VUNPCKLPD128) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vunpcklpd128(&vals1, &vals2, &ret)

	log.Printf("VUNPCKLPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VUNPCKLPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
