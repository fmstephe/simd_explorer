package pxor

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpxor_128.s
var assemblyVpxor128 string

//go:embed stub_vpxor_128.go
var stubVpxor128 string

type VPXOR128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPXOR128() *VPXOR128 {
	return &VPXOR128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPXOR128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPXOR128) Output() *number.Parameter {
	return v.ret
}

func (v *VPXOR128) Name() string {
	return "VPXOR (128 bit) "
}

func (v *VPXOR128) Description() string {
	return "Bitwise XOR of packed bytes."
}

func (v *VPXOR128) Stub() string {
	return stubVpxor128
}

func (v *VPXOR128) Assembly() string {
	return assemblyVpxor128
}

func (v *VPXOR128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpxor128(&vals1, &vals2, &ret)

	log.Printf("VPXOR128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPXOR128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
