package packss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpackssdw_256.s
var assemblyVpackssdw256 string

//go:embed stub_vpackssdw_256.go
var stubVpackssdw256 string

type VPACKSSDW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKSSDW256() *VPACKSSDW256 {
	return &VPACKSSDW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPACKSSDW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKSSDW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKSSDW256) Name() string {
	return "VPACKSSDW (256 bit) "
}

func (v *VPACKSSDW256) Description() string {
	return "Pack signed 32-bit integers from two sources into signed 16-bit with saturation."
}

func (v *VPACKSSDW256) Stub() string {
	return stubVpackssdw256
}

func (v *VPACKSSDW256) Assembly() string {
	return assemblyVpackssdw256
}

func (v *VPACKSSDW256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [16]int16{}

	vpackssdw256(&vals1, &vals2, &ret)

	log.Printf("VPACKSSDW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPACKSSDW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
