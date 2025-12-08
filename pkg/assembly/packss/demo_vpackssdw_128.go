package packss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackssdw_128.s
var assemblyVpackssdw128 string

//go:embed stub_vpackssdw_128.go
var stubVpackssdw128 string

type VPACKSSDW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKSSDW128() *VPACKSSDW128 {
	return &VPACKSSDW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPACKSSDW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKSSDW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKSSDW128) Name() string {
	return "VPACKSSDW (128 bit) "
}

func (v *VPACKSSDW128) Description() string {
	return "Pack signed 32-bit integers from two sources into signed 16-bit with saturation."
}

func (v *VPACKSSDW128) Stub() string {
	return stubVpackssdw128
}

func (v *VPACKSSDW128) Assembly() string {
	return assemblyVpackssdw128
}

func (v *VPACKSSDW128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	vpackssdw128(&vals1, &vals2, &ret)

	log.Printf("VPACKSSDW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPACKSSDW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
