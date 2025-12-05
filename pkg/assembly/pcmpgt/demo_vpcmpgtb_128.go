package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtb_128.s
var assemblyVpcmpgtb128 string

//go:embed stub_vpcmpgtb_128.go
var stubVpcmpgtb128 string

type VPCMPGTB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTB128() *VPCMPGTB128 {
	return &VPCMPGTB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPCMPGTB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTB128) Name() string {
	return "VPCMPGTB (128 bit) "
}

func (v *VPCMPGTB128) Description() string {
	return "Compare packed signed 8-bit integers for greater-than; bytes are 0xFF if vals1 > vals2 else 0x00."
}

func (v *VPCMPGTB128) Stub() string {
	return stubVpcmpgtb128
}

func (v *VPCMPGTB128) Assembly() string {
	return assemblyVpcmpgtb128
}

func (v *VPCMPGTB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpcmpgtb128(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPCMPGTB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
