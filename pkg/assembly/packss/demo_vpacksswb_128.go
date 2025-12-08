package packss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpacksswb_128.s
var assemblyVpacksswb128 string

//go:embed stub_vpacksswb_128.go
var stubVpacksswb128 string

type VPACKSSWB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPACKSSWB128() *VPACKSSWB128 {
	return &VPACKSSWB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 8, 10),
	}
}

func (v *VPACKSSWB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPACKSSWB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPACKSSWB128) Name() string {
	return "VPACKSSWB (128 bit) "
}

func (v *VPACKSSWB128) Description() string {
	return "Pack signed 16-bit integers from two sources into signed 8-bit with saturation."
}

func (v *VPACKSSWB128) Stub() string {
	return stubVpacksswb128
}

func (v *VPACKSSWB128) Assembly() string {
	return assemblyVpacksswb128
}

func (v *VPACKSSWB128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [16]int8{}

	vpacksswb128(&vals1, &vals2, &ret)

	log.Printf("VPACKSSWB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Int8SliceToBytes(ret[:]))
}

func (v *VPACKSSWB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
