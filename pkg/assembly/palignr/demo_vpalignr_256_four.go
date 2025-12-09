package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_four.s
var assemblyVpalignr256Four string

//go:embed stub_vpalignr_256_four.go
var stubVpalignr256Four string

type VPALIGNR256FOUR struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256FOUR() *VPALIGNR256FOUR {
	return &VPALIGNR256FOUR{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256FOUR) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256FOUR) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256FOUR) Name() string {
	return "VPALIGNR (256 bit) four"
}

func (v *VPALIGNR256FOUR) Description() string {
	return "Align right by 4 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256FOUR) Stub() string {
	return stubVpalignr256Four
}

func (v *VPALIGNR256FOUR) Assembly() string {
	return assemblyVpalignr256Four
}

func (v *VPALIGNR256FOUR) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Four(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256FOUR vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256FOUR) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
