package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqb_128.s
var assemblyVpcmpeqb128 string

//go:embed stub_vpcmpeqb_128.go
var stubVpcmpeqb128 string

type VPCMPEQB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQB128() *VPCMPEQB128 {
	return &VPCMPEQB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPCMPEQB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQB128) Name() string {
	return "VPCMPEQB (128 bit) "
}

func (v *VPCMPEQB128) Description() string {
	return "Compare packed 8-bit integers for equality; result bytes are 0xFF if equal, else 0x00."
}

func (v *VPCMPEQB128) Stub() string {
	return stubVpcmpeqb128
}

func (v *VPCMPEQB128) Assembly() string {
	return assemblyVpcmpeqb128
}

func (v *VPCMPEQB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpcmpeqb128(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPCMPEQB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
