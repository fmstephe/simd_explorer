package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignd_128.s
var assemblyVpsignd128 string

//go:embed stub_vpsignd_128.go
var stubVpsignd128 string

type VPSIGND128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGND128() *VPSIGND128 {
	return &VPSIGND128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPSIGND128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGND128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGND128) Name() string {
	return "VPSIGND (128 bit) "
}

func (v *VPSIGND128) Description() string {
	return "Apply signs from vals2 to vals1 (signed 32-bit). ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]). Zero in vals2 yields zero."
}

func (v *VPSIGND128) Stub() string {
	return stubVpsignd128
}

func (v *VPSIGND128) Assembly() string {
	return assemblyVpsignd128
}

func (v *VPSIGND128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpsignd128(&vals1, &vals2, &ret)

	log.Printf("VPSIGND vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGND128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
