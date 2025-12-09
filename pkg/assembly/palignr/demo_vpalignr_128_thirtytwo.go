package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_thirtytwo.s
var assemblyVpalignr128Thirtytwo string

//go:embed stub_vpalignr_128_thirtytwo.go
var stubVpalignr128Thirtytwo string

type VPALIGNR128THIRTYTWO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128THIRTYTWO() *VPALIGNR128THIRTYTWO {
	return &VPALIGNR128THIRTYTWO{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128THIRTYTWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128THIRTYTWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128THIRTYTWO) Name() string {
	return "VPALIGNR (128 bit) thirtytwo"
}

func (v *VPALIGNR128THIRTYTWO) Description() string {
	return "Align right by 32 bytes across vals1 and vals2."
}

func (v *VPALIGNR128THIRTYTWO) Stub() string {
	return stubVpalignr128Thirtytwo
}

func (v *VPALIGNR128THIRTYTWO) Assembly() string {
	return assemblyVpalignr128Thirtytwo
}

func (v *VPALIGNR128THIRTYTWO) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Thirtytwo(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128THIRTYTWO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128THIRTYTWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
