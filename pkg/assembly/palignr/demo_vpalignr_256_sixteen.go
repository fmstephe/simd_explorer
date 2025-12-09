package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_sixteen.s
var assemblyVpalignr256Sixteen string

//go:embed stub_vpalignr_256_sixteen.go
var stubVpalignr256Sixteen string

type VPALIGNR256SIXTEEN struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256SIXTEEN() *VPALIGNR256SIXTEEN {
	return &VPALIGNR256SIXTEEN{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256SIXTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256SIXTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256SIXTEEN) Name() string {
	return "VPALIGNR (256 bit) sixteen"
}

func (v *VPALIGNR256SIXTEEN) Description() string {
	return "Align right by 16 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256SIXTEEN) Stub() string {
	return stubVpalignr256Sixteen
}

func (v *VPALIGNR256SIXTEEN) Assembly() string {
	return assemblyVpalignr256Sixteen
}

func (v *VPALIGNR256SIXTEEN) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Sixteen(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256SIXTEEN vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256SIXTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
