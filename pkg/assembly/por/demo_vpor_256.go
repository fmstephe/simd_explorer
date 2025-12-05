package por

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpor_256.s
var assemblyVpor256 string

//go:embed stub_vpor_256.go
var stubVpor256 string

type VPOR256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPOR256() *VPOR256 {
	return &VPOR256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPOR256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPOR256) Output() *number.Parameter {
	return v.ret
}

func (v *VPOR256) Name() string {
	return "VPOR (256 bit) "
}

func (v *VPOR256) Description() string {
	return "Bitwise OR of packed bytes."
}

func (v *VPOR256) Stub() string {
	return stubVpor256
}

func (v *VPOR256) Assembly() string {
	return assemblyVpor256
}

func (v *VPOR256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpor256(&vals1, &vals2, &ret)

	log.Printf("VPOR256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPOR256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
