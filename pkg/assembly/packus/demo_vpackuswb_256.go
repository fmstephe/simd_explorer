package packus

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackuswb_256.s
var assemblyVpackuswb256 string

//go:embed stub_vpackuswb_256.go
var stubVpackuswb256 string

type VPACKUSWB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKUSWB256() *VPACKUSWB256 {
	return &VPACKUSWB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPACKUSWB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKUSWB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKUSWB256) Name() string {
	return "VPACKUSWB (256 bit) "
}

func (v *VPACKUSWB256) Description() string {
	return "Pack signed 16-bit integers from two sources into unsigned 8-bit with saturation."
}

func (v *VPACKUSWB256) Stub() string {
	return stubVpackuswb256
}

func (v *VPACKUSWB256) Assembly() string {
	return assemblyVpackuswb256
}

func (v *VPACKUSWB256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [32]int8{}

	vpackuswb256(&vals1, &vals2, &ret)

	log.Printf("VPACKUSWB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int8SliceToBytes(ret[:]))
}

func (v *VPACKUSWB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
