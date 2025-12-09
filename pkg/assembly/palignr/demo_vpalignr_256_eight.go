package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_eight.s
var assemblyVpalignr256Eight string

//go:embed stub_vpalignr_256_eight.go
var stubVpalignr256Eight string

type VPALIGNR256EIGHT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256EIGHT() *VPALIGNR256EIGHT {
	return &VPALIGNR256EIGHT{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256EIGHT) Name() string {
	return "VPALIGNR (256 bit) eight"
}

func (v *VPALIGNR256EIGHT) Description() string {
	return "Align right by 8 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256EIGHT) Stub() string {
	return stubVpalignr256Eight
}

func (v *VPALIGNR256EIGHT) Assembly() string {
	return assemblyVpalignr256Eight
}

func (v *VPALIGNR256EIGHT) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Eight(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256EIGHT vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
