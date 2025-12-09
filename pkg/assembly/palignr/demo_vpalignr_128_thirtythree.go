package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_thirtythree.s
var assemblyVpalignr128Thirtythree string

//go:embed stub_vpalignr_128_thirtythree.go
var stubVpalignr128Thirtythree string

type VPALIGNR128THIRTYTHREE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128THIRTYTHREE() *VPALIGNR128THIRTYTHREE {
	return &VPALIGNR128THIRTYTHREE{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128THIRTYTHREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128THIRTYTHREE) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128THIRTYTHREE) Name() string {
	return "VPALIGNR (128 bit) thirtythree"
}

func (v *VPALIGNR128THIRTYTHREE) Description() string {
	return "Align right by 33 bytes across vals1 and vals2."
}

func (v *VPALIGNR128THIRTYTHREE) Stub() string {
	return stubVpalignr128Thirtythree
}

func (v *VPALIGNR128THIRTYTHREE) Assembly() string {
	return assemblyVpalignr128Thirtythree
}

func (v *VPALIGNR128THIRTYTHREE) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Thirtythree(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128THIRTYTHREE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128THIRTYTHREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
