package packus

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackusdw_128.s
var assemblyVpackusdw128 string

//go:embed stub_vpackusdw_128.go
var stubVpackusdw128 string

type VPACKUSDW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKUSDW128() *VPACKUSDW128 {
	return &VPACKUSDW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPACKUSDW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKUSDW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKUSDW128) Name() string {
	return "VPACKUSDW (128 bit) "
}

func (v *VPACKUSDW128) Description() string {
	return "Pack signed 32-bit integers from two sources into unsigned 16-bit with saturation."
}

func (v *VPACKUSDW128) Stub() string {
	return stubVpackusdw128
}

func (v *VPACKUSDW128) Assembly() string {
	return assemblyVpackusdw128
}

func (v *VPACKUSDW128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	vpackusdw128(&vals1, &vals2, &ret)

	log.Printf("VPACKUSDW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int16SliceToBytes(ret[:]))
}

func (v *VPACKUSDW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
