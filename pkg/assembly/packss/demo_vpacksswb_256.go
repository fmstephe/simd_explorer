package packss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpacksswb_256.s
var assemblyVpacksswb256 string

//go:embed stub_vpacksswb_256.go
var stubVpacksswb256 string

type VPACKSSWB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKSSWB256() *VPACKSSWB256 {
	return &VPACKSSWB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 8, 10),
	}
}

func (v *VPACKSSWB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKSSWB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKSSWB256) Name() string {
	return "VPACKSSWB (256 bit) "
}

func (v *VPACKSSWB256) Description() string {
	return "Pack signed 16-bit integers from two sources into signed 8-bit with saturation."
}

func (v *VPACKSSWB256) Stub() string {
	return stubVpacksswb256
}

func (v *VPACKSSWB256) Assembly() string {
	return assemblyVpacksswb256
}

func (v *VPACKSSWB256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [32]int8{}

	vpacksswb256(&vals1, &vals2, &ret)

	log.Printf("VPACKSSWB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int8SliceToBytes(ret[:]))
}

func (v *VPACKSSWB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
