package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignb_256.s
var assemblyVpsignb256 string

//go:embed stub_vpsignb_256.go
var stubVpsignb256 string

type VPSIGNB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGNB256() *VPSIGNB256 {
	return &VPSIGNB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 8, 10),
	}
}

func (v *VPSIGNB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGNB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGNB256) Name() string {
	return "VPSIGNB (256 bit) "
}

func (v *VPSIGNB256) Description() string {
	return "Apply signs from vals2 to vals1 (signed 8-bit), per 128-bit lane. " +
		"ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]); zero in vals2 yields zero."
}

func (v *VPSIGNB256) Stub() string {
	return stubVpsignb256
}

func (v *VPSIGNB256) Assembly() string {
	return assemblyVpsignb256
}

func (v *VPSIGNB256) Run() {
	vals1 := [32]int8{}
	copy(vals1[:], number.ToInt8Slice(v.vals1.FlatData()))
	vals2 := [32]int8{}
	copy(vals2[:], number.ToInt8Slice(v.vals2.FlatData()))
	ret := [32]int8{}
	copy(ret[:], number.ToInt8Slice(v.ret.FlatData()))

	vpsignb256(&vals1, &vals2, &ret)

	log.Printf("VPSIGNB vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGNB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
