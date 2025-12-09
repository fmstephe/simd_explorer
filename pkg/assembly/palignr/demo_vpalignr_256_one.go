package palignr

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpalignr_256_one.s
var assemblyVpalignr256One string

//go:embed stub_vpalignr_256_one.go
var stubVpalignr256One string

type VPALIGNR256ONE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPALIGNR256ONE() *VPALIGNR256ONE {
	return &VPALIGNR256ONE{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPALIGNR256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPALIGNR256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPALIGNR256ONE) Name() string {
	return "VPALIGNR (256 bit) one"
}

func (v *VPALIGNR256ONE) Description() string {
	return "Align right by 1 byte across vals1 and vals2 per 128-bit lane."
}

func (v *VPALIGNR256ONE) Stub() string {
	return stubVpalignr256One
}

func (v *VPALIGNR256ONE) Assembly() string {
	return assemblyVpalignr256One
}

func (v *VPALIGNR256ONE) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpalignr256One(&vals1, &vals2, &ret)

	log.Printf("VPALIGNR256ONE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPALIGNR256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
