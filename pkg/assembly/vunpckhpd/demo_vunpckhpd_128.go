package vunpckhpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhpd_128.s
var assemblyVunpckhpd128 string

//go:embed stub_vunpckhpd_128.go
var stubVunpckhpd128 string

type VUNPCKHPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKHPD128() *VUNPCKHPD128 {
	return &VUNPCKHPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VUNPCKHPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKHPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKHPD128) Name() string {
	return "VUNPCKHPD (128 bit) "
}

func (v *VUNPCKHPD128) Description() string {
	return "Unpack high doubles: ret = [vals1[1], vals2[1]]."
}

func (v *VUNPCKHPD128) Stub() string {
	return stubVunpckhpd128
}

func (v *VUNPCKHPD128) Assembly() string {
	return assemblyVunpckhpd128
}

func (v *VUNPCKHPD128) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vunpckhpd128(&vals1, &vals2, &ret)

	log.Printf("VUNPCKHPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VUNPCKHPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
