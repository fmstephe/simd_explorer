package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_three.s
var assemblyVpalignr128Three string

//go:embed stub_vpalignr_128_three.go
var stubVpalignr128Three string

type VPALIGNR128THREE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128THREE() *VPALIGNR128THREE {
	return &VPALIGNR128THREE{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128THREE) Name() string {
	return "VPALIGNR (128 bit) three"
}

func (v *VPALIGNR128THREE) Description() string {
	return "Align right by 3 bytes across vals1 and vals2."
}

func (v *VPALIGNR128THREE) Stub() string {
	return stubVpalignr128Three
}

func (v *VPALIGNR128THREE) Assembly() string {
	return assemblyVpalignr128Three
}

func (v *VPALIGNR128THREE) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Three(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128THREE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
