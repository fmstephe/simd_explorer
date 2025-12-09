package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_zero.s
var assemblyVpalignr256Zero string

//go:embed stub_vpalignr_256_zero.go
var stubVpalignr256Zero string

type VPALIGNR256ZERO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256ZERO() *VPALIGNR256ZERO {
	return &VPALIGNR256ZERO{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256ZERO) Name() string {
	return "VPALIGNR (256 bit) zero"
}

func (v *VPALIGNR256ZERO) Description() string {
	return "Align right by 0 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256ZERO) Stub() string {
	return stubVpalignr256Zero
}

func (v *VPALIGNR256ZERO) Assembly() string {
	return assemblyVpalignr256Zero
}

func (v *VPALIGNR256ZERO) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Zero(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256ZERO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
