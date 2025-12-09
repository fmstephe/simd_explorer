package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_zero.s
var assemblyVpalignr128Zero string

//go:embed stub_vpalignr_128_zero.go
var stubVpalignr128Zero string

type VPALIGNR128ZERO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128ZERO() *VPALIGNR128ZERO {
	return &VPALIGNR128ZERO{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128ZERO) Name() string {
	return "VPALIGNR (128 bit) zero"
}

func (v *VPALIGNR128ZERO) Description() string {
	return "Align right by 0 bytes across vals1 and vals2; per 128-bit lane."
}

func (v *VPALIGNR128ZERO) Stub() string {
	return stubVpalignr128Zero
}

func (v *VPALIGNR128ZERO) Assembly() string {
	return assemblyVpalignr128Zero
}

func (v *VPALIGNR128ZERO) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Zero(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128ZERO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
