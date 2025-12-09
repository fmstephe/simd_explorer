package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_thirtytwo.s
var assemblyVpalignr256Thirtytwo string

//go:embed stub_vpalignr_256_thirtytwo.go
var stubVpalignr256Thirtytwo string

type VPALIGNR256THIRTYTWO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256THIRTYTWO() *VPALIGNR256THIRTYTWO {
	return &VPALIGNR256THIRTYTWO{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256THIRTYTWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256THIRTYTWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256THIRTYTWO) Name() string {
	return "VPALIGNR (256 bit) thirtytwo"
}

func (v *VPALIGNR256THIRTYTWO) Description() string {
	return "Align right by 32 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256THIRTYTWO) Stub() string {
	return stubVpalignr256Thirtytwo
}

func (v *VPALIGNR256THIRTYTWO) Assembly() string {
	return assemblyVpalignr256Thirtytwo
}

func (v *VPALIGNR256THIRTYTWO) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Thirtytwo(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256THIRTYTWO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256THIRTYTWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
