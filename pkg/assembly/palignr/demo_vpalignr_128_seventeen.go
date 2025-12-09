package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_seventeen.s
var assemblyVpalignr128Seventeen string

//go:embed stub_vpalignr_128_seventeen.go
var stubVpalignr128Seventeen string

type VPALIGNR128SEVENTEEN struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128SEVENTEEN() *VPALIGNR128SEVENTEEN {
	return &VPALIGNR128SEVENTEEN{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128SEVENTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128SEVENTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128SEVENTEEN) Name() string {
	return "VPALIGNR (128 bit) seventeen"
}

func (v *VPALIGNR128SEVENTEEN) Description() string {
	return "Align right by 17 bytes across vals1 and vals2."
}

func (v *VPALIGNR128SEVENTEEN) Stub() string {
	return stubVpalignr128Seventeen
}

func (v *VPALIGNR128SEVENTEEN) Assembly() string {
	return assemblyVpalignr128Seventeen
}

func (v *VPALIGNR128SEVENTEEN) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Seventeen(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128SEVENTEEN vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128SEVENTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
