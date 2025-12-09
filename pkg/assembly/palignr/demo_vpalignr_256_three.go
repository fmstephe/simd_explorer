package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_three.s
var assemblyVpalignr256Three string

//go:embed stub_vpalignr_256_three.go
var stubVpalignr256Three string

type VPALIGNR256THREE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256THREE() *VPALIGNR256THREE {
	return &VPALIGNR256THREE{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256THREE) Name() string {
	return "VPALIGNR (256 bit) three"
}

func (v *VPALIGNR256THREE) Description() string {
	return "Align right by 3 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256THREE) Stub() string {
	return stubVpalignr256Three
}

func (v *VPALIGNR256THREE) Assembly() string {
	return assemblyVpalignr256Three
}

func (v *VPALIGNR256THREE) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Three(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256THREE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
