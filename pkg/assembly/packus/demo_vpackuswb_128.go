package packus

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackuswb_128.s
var assemblyVpackuswb128 string

//go:embed stub_vpackuswb_128.go
var stubVpackuswb128 string

type VPACKUSWB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKUSWB128() *VPACKUSWB128 {
	return &VPACKUSWB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPACKUSWB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKUSWB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKUSWB128) Name() string {
	return "VPACKUSWB (128 bit) "
}

func (v *VPACKUSWB128) Description() string {
	return "Pack signed 16-bit integers from two sources into unsigned 8-bit with saturation."
}

func (v *VPACKUSWB128) Stub() string {
	return stubVpackuswb128
}

func (v *VPACKUSWB128) Assembly() string {
	return assemblyVpackuswb128
}

func (v *VPACKUSWB128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [16]int8{}

	vpackuswb128(&vals1, &vals2, &ret)

	log.Printf("VPACKUSWB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int8SliceToBytes(ret[:]))
}

func (v *VPACKUSWB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
