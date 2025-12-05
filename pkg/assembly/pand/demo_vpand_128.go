package pand

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpand_128.s
var assemblyVpand128 string

//go:embed stub_vpand_128.go
var stubVpand128 string

type VPAND128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAND128() *VPAND128 {
	return &VPAND128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPAND128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAND128) Output() *number.Parameter {
	return v.ret
}

func (v *VPAND128) Name() string {
	return "VPAND (128 bit) "
}

func (v *VPAND128) Description() string {
	return "Bitwise AND of packed bytes."
}

func (v *VPAND128) Stub() string {
	return stubVpand128
}

func (v *VPAND128) Assembly() string {
	return assemblyVpand128
}

func (v *VPAND128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpand128(&vals1, &vals2, &ret)

	log.Printf("VPAND128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPAND128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
