package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqw_256.s
var assemblyVpcmpeqw256 string

//go:embed stub_vpcmpeqw_256.go
var stubVpcmpeqw256 string

type VPCMPEQW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQW256() *VPCMPEQW256 {
	return &VPCMPEQW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 16),
	}
}

func (v *VPCMPEQW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQW256) Name() string {
	return "VPCMPEQW (256 bit) "
}

func (v *VPCMPEQW256) Description() string {
	return "Compare packed 16-bit integers for equality; result words are 0xFFFF if equal, else 0x0000."
}

func (v *VPCMPEQW256) Stub() string {
	return stubVpcmpeqw256
}

func (v *VPCMPEQW256) Assembly() string {
	return assemblyVpcmpeqw256
}

func (v *VPCMPEQW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpcmpeqw256(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
