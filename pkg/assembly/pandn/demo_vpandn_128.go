package pandn

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpandn_128.s
var assemblyVpandn128 string

//go:embed stub_vpandn_128.go
var stubVpandn128 string

type VPANDN128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPANDN128() *VPANDN128 {
	return &VPANDN128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPANDN128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPANDN128) Output() *number.Parameter {
	return v.ret
}

func (v *VPANDN128) Name() string {
	return "VPANDN (128 bit) "
}

func (v *VPANDN128) Description() string {
	return "Bitwise AND NOT of packed bytes: (~vals1) AND vals2."
}

func (v *VPANDN128) Stub() string {
	return stubVpandn128
}

func (v *VPANDN128) Assembly() string {
	return assemblyVpandn128
}

func (v *VPANDN128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpandn128(&vals1, &vals2, &ret)

	log.Printf("VPANDN128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPANDN128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
