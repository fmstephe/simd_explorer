package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_128_two.s
var assemblyVpalignr128Two string

//go:embed stub_vpalignr_128_two.go
var stubVpalignr128Two string

type VPALIGNR128TWO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR128TWO() *VPALIGNR128TWO {
	return &VPALIGNR128TWO{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPALIGNR128TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR128TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR128TWO) Name() string {
	return "VPALIGNR (128 bit) two"
}

func (v *VPALIGNR128TWO) Description() string {
	return "Align right by 2 bytes across vals1 and vals2."
}

func (v *VPALIGNR128TWO) Stub() string {
	return stubVpalignr128Two
}

func (v *VPALIGNR128TWO) Assembly() string {
	return assemblyVpalignr128Two
}

func (v *VPALIGNR128TWO) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpalignr128Two(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR128TWO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR128TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
