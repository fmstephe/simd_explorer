package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_seventeen.s
var assemblyVpalignr256Seventeen string

//go:embed stub_vpalignr_256_seventeen.go
var stubVpalignr256Seventeen string

type VPALIGNR256SEVENTEEN struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256SEVENTEEN() *VPALIGNR256SEVENTEEN {
	return &VPALIGNR256SEVENTEEN{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256SEVENTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256SEVENTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256SEVENTEEN) Name() string {
	return "VPALIGNR (256 bit) seventeen"
}

func (v *VPALIGNR256SEVENTEEN) Description() string {
	return "Align right by 17 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256SEVENTEEN) Stub() string {
	return stubVpalignr256Seventeen
}

func (v *VPALIGNR256SEVENTEEN) Assembly() string {
	return assemblyVpalignr256Seventeen
}

func (v *VPALIGNR256SEVENTEEN) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Seventeen(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256SEVENTEEN vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256SEVENTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
