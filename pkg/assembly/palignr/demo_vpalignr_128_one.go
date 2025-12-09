package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_one.s
var assemblyVpalignr128One string

//go:embed stub_vpalignr_128_one.go
var stubVpalignr128One string

type VPALIGNR128ONE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128ONE() *VPALIGNR128ONE {
	return &VPALIGNR128ONE{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128ONE) Name() string {
	return "VPALIGNR (128 bit) one"
}

func (v *VPALIGNR128ONE) Description() string {
	return "Align right by 1 byte across vals1 and vals2."
}

func (v *VPALIGNR128ONE) Stub() string {
	return stubVpalignr128One
}

func (v *VPALIGNR128ONE) Assembly() string {
	return assemblyVpalignr128One
}

func (v *VPALIGNR128ONE) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128One(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128ONE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
