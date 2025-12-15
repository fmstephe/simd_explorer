package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignw_256.s
var assemblyVpsignw256 string

//go:embed stub_vpsignw_256.go
var stubVpsignw256 string

type VPSIGNW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGNW256() *VPSIGNW256 {
	return &VPSIGNW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPSIGNW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGNW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGNW256) Name() string {
	return "VPSIGNW (256 bit) "
}

func (v *VPSIGNW256) Description() string {
	return "Apply signs from vals2 to vals1 (signed 16-bit), per 128-bit lane. " +
		"ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]); zero in vals2 yields zero."
}

func (v *VPSIGNW256) Stub() string {
	return stubVpsignw256
}

func (v *VPSIGNW256) Assembly() string {
	return assemblyVpsignw256
}

func (v *VPSIGNW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [16]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpsignw256(&vals1, &vals2, &ret)

	log.Printf("VPSIGNW vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGNW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
