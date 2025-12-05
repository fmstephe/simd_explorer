package pand

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpand_256.s
var assemblyVpand256 string

//go:embed stub_vpand_256.go
var stubVpand256 string

type VPAND256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAND256() *VPAND256 {
	return &VPAND256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPAND256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAND256) Output() *number.Parameter {
	return v.ret
}

func (v *VPAND256) Name() string {
	return "VPAND (256 bit) "
}

func (v *VPAND256) Description() string {
	return "Bitwise AND of packed bytes."
}

func (v *VPAND256) Stub() string {
	return stubVpand256
}

func (v *VPAND256) Assembly() string {
	return assemblyVpand256
}

func (v *VPAND256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpand256(&vals1, &vals2, &ret)

	log.Printf("VPAND256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPAND256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
