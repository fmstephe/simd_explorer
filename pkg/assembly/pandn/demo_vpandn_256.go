package pandn

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpandn_256.s
var assemblyVpandn256 string

//go:embed stub_vpandn_256.go
var stubVpandn256 string

type VPANDN256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPANDN256() *VPANDN256 {
	return &VPANDN256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPANDN256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPANDN256) Output() *number.Parameter {
	return v.ret
}

func (v *VPANDN256) Name() string {
	return "VPANDN (256 bit) "
}

func (v *VPANDN256) Description() string {
	return "Bitwise AND NOT of packed bytes: (~vals1) AND vals2."
}

func (v *VPANDN256) Stub() string {
	return stubVpandn256
}

func (v *VPANDN256) Assembly() string {
	return assemblyVpandn256
}

func (v *VPANDN256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpandn256(&vals1, &vals2, &ret)

	log.Printf("VPANDN256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPANDN256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
