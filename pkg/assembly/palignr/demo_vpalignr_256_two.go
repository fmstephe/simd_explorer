package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_two.s
var assemblyVpalignr256Two string

//go:embed stub_vpalignr_256_two.go
var stubVpalignr256Two string

type VPALIGNR256TWO struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256TWO() *VPALIGNR256TWO {
	return &VPALIGNR256TWO{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256TWO) Name() string {
	return "VPALIGNR (256 bit) two"
}

func (v *VPALIGNR256TWO) Description() string {
	return "Align right by 2 bytes across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256TWO) Stub() string {
	return stubVpalignr256Two
}

func (v *VPALIGNR256TWO) Assembly() string {
	return assemblyVpalignr256Two
}

func (v *VPALIGNR256TWO) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256Two(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256TWO vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
