package vpsign

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsignd_256.s
var assemblyVpsignd256 string

//go:embed stub_vpsignd_256.go
var stubVpsignd256 string

type VPSIGND256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSIGND256() *VPSIGND256 {
	return &VPSIGND256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPSIGND256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSIGND256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSIGND256) Name() string {
	return "VPSIGND (256 bit) "
}

func (v *VPSIGND256) Description() string {
	return "Apply signs from vals2 to vals1 (signed 32-bit), per 128-bit lane. ret[i] = abs(vals1[i]) if vals2[i] >= 0, else -abs(vals1[i]); zero in vals2 yields zero."
}

func (v *VPSIGND256) Stub() string {
	return stubVpsignd256
}

func (v *VPSIGND256) Assembly() string {
	return assemblyVpsignd256
}

func (v *VPSIGND256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpsignd256(&vals1, &vals2, &ret)

	log.Printf("VPSIGND vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSIGND256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
