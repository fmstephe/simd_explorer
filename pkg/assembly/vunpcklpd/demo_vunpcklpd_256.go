package vunpcklpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklpd_256.s
var assemblyVunpcklpd256 string

//go:embed stub_vunpcklpd_256.go
var stubVunpcklpd256 string

type VUNPCKLPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKLPD256() *VUNPCKLPD256 {
	return &VUNPCKLPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VUNPCKLPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKLPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKLPD256) Name() string {
	return "VUNPCKLPD (256 bit) "
}

func (v *VUNPCKLPD256) Description() string {
	return "Unpack low doubles per 128-bit lane: ret = [vals1[0], vals2[0], vals1[2], vals2[2]]."
}

func (v *VUNPCKLPD256) Stub() string {
	return stubVunpcklpd256
}

func (v *VUNPCKLPD256) Assembly() string {
	return assemblyVunpcklpd256
}

func (v *VUNPCKLPD256) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vunpcklpd256(&vals1, &vals2, &ret)

	log.Printf("VUNPCKLPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VUNPCKLPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
