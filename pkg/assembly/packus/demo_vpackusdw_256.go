package packus

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackusdw_256.s
var assemblyVpackusdw256 string

//go:embed stub_vpackusdw_256.go
var stubVpackusdw256 string

type VPACKUSDW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKUSDW256() *VPACKUSDW256 {
	return &VPACKUSDW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPACKUSDW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKUSDW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKUSDW256) Name() string {
	return "VPACKUSDW (256 bit) "
}

func (v *VPACKUSDW256) Description() string {
	return "Pack signed 32-bit integers from two sources into unsigned 16-bit with saturation."
}

func (v *VPACKUSDW256) Stub() string {
	return stubVpackusdw256
}

func (v *VPACKUSDW256) Assembly() string {
	return assemblyVpackusdw256
}

func (v *VPACKUSDW256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [16]int16{}

	vpackusdw256(&vals1, &vals2, &ret)

	log.Printf("VPACKUSDW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int16SliceToBytes(ret[:]))
}

func (v *VPACKUSDW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
