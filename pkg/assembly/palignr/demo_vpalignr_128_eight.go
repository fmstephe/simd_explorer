package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_eight.s
var assemblyVpalignr128Eight string

//go:embed stub_vpalignr_128_eight.go
var stubVpalignr128Eight string

type VPALIGNR128EIGHT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128EIGHT() *VPALIGNR128EIGHT {
	return &VPALIGNR128EIGHT{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128EIGHT) Name() string {
	return "VPALIGNR (128 bit) eight"
}

func (v *VPALIGNR128EIGHT) Description() string {
	return "Align right by 8 bytes across vals1 and vals2."
}

func (v *VPALIGNR128EIGHT) Stub() string {
	return stubVpalignr128Eight
}

func (v *VPALIGNR128EIGHT) Assembly() string {
	return assemblyVpalignr128Eight
}

func (v *VPALIGNR128EIGHT) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Eight(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128EIGHT vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
