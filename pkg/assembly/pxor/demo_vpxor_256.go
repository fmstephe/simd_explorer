package pxor

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpxor_256.s
var assemblyVpxor256 string

//go:embed stub_vpxor_256.go
var stubVpxor256 string

type VPXOR256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPXOR256() *VPXOR256 {
	return &VPXOR256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPXOR256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPXOR256) Output() *number.Parameter {
	return v.ret
}

func (v *VPXOR256) Name() string {
	return "VPXOR (256 bit) "
}

func (v *VPXOR256) Description() string {
	return "Bitwise XOR of packed bytes."
}

func (v *VPXOR256) Stub() string {
	return stubVpxor256
}

func (v *VPXOR256) Assembly() string {
	return assemblyVpxor256
}

func (v *VPXOR256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpxor256(&vals1, &vals2, &ret)

	log.Printf("VPXOR256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPXOR256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
