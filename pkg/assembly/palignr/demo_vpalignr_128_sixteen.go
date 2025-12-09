package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_sixteen.s
var assemblyVpalignr128Sixteen string

//go:embed stub_vpalignr_128_sixteen.go
var stubVpalignr128Sixteen string

type VPALIGNR128SIXTEEN struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128SIXTEEN() *VPALIGNR128SIXTEEN {
	return &VPALIGNR128SIXTEEN{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128SIXTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128SIXTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128SIXTEEN) Name() string {
	return "VPALIGNR (128 bit) sixteen"
}

func (v *VPALIGNR128SIXTEEN) Description() string {
	return "Align right by 16 bytes across vals1 and vals2."
}

func (v *VPALIGNR128SIXTEEN) Stub() string {
	return stubVpalignr128Sixteen
}

func (v *VPALIGNR128SIXTEEN) Assembly() string {
	return assemblyVpalignr128Sixteen
}

func (v *VPALIGNR128SIXTEEN) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Sixteen(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128SIXTEEN vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128SIXTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
