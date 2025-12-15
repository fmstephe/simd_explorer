package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignb_128.s
var assemblyVpsignb128 string

//go:embed stub_vpsignb_128.go
var stubVpsignb128 string

type VPSIGNB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGNB128() *VPSIGNB128 {
	return &VPSIGNB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 8, 10),
	}
}

func (v *VPSIGNB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGNB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGNB128) Name() string {
	return "VPSIGNB (128 bit) "
}

func (v *VPSIGNB128) Description() string {
	return "Apply signs from vals2 to vals1 (signed 8-bit). " +
		"ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]). Zero in vals2 yields zero."
}

func (v *VPSIGNB128) Stub() string {
	return stubVpsignb128
}

func (v *VPSIGNB128) Assembly() string {
	return assemblyVpsignb128
}

func (v *VPSIGNB128) Run() {
	vals1 := [16]int8{}
	copy(vals1[:], number.ToInt8Slice(v.vals1.FlatData()))
	vals2 := [16]int8{}
	copy(vals2[:], number.ToInt8Slice(v.vals2.FlatData()))
	ret := [16]int8{}
	copy(ret[:], number.ToInt8Slice(v.ret.FlatData()))

	vpsignb128(&vals1, &vals2, &ret)

	log.Printf("VPSIGNB vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGNB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
