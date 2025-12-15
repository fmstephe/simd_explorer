package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignw_128.s
var assemblyVpsignw128 string

//go:embed stub_vpsignw_128.go
var stubVpsignw128 string

type VPSIGNW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGNW128() *VPSIGNW128 {
	return &VPSIGNW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPSIGNW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGNW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGNW128) Name() string {
	return "VPSIGNW (128 bit) "
}

func (v *VPSIGNW128) Description() string {
	return "Apply signs from vals2 to vals1 (signed 16-bit). " +
		"ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]). Zero in vals2 yields zero."
}

func (v *VPSIGNW128) Stub() string {
	return stubVpsignw128
}

func (v *VPSIGNW128) Assembly() string {
	return assemblyVpsignw128
}

func (v *VPSIGNW128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [8]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpsignw128(&vals1, &vals2, &ret)

	log.Printf("VPSIGNW vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGNW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
