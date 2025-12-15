package vunpckhpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhpd_256.s
var assemblyVunpckhpd256 string

//go:embed stub_vunpckhpd_256.go
var stubVunpckhpd256 string

type VUNPCKHPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKHPD256() *VUNPCKHPD256 {
	return &VUNPCKHPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VUNPCKHPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKHPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKHPD256) Name() string {
	return "VUNPCKHPD (256 bit) "
}

func (v *VUNPCKHPD256) Description() string {
	return "Unpack high doubles per 128-bit lane: ret = [vals1[1], vals2[1], vals1[3], vals2[3]]."
}

func (v *VUNPCKHPD256) Stub() string {
	return stubVunpckhpd256
}

func (v *VUNPCKHPD256) Assembly() string {
	return assemblyVunpckhpd256
}

func (v *VUNPCKHPD256) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vunpckhpd256(&vals1, &vals2, &ret)

	log.Printf("VUNPCKHPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VUNPCKHPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
