package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqb_256.s
var assemblyVpcmpeqb256 string

//go:embed stub_vpcmpeqb_256.go
var stubVpcmpeqb256 string

type VPCMPEQB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQB256() *VPCMPEQB256 {
	return &VPCMPEQB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPCMPEQB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQB256) Name() string {
	return "VPCMPEQB (256 bit) "
}

func (v *VPCMPEQB256) Description() string {
	return "Compare packed 8-bit integers for equality; result bytes are 0xFF if equal, else 0x00."
}

func (v *VPCMPEQB256) Stub() string {
	return stubVpcmpeqb256
}

func (v *VPCMPEQB256) Assembly() string {
	return assemblyVpcmpeqb256
}

func (v *VPCMPEQB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpcmpeqb256(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPCMPEQB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
