package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqw_128.s
var assemblyVpcmpeqw128 string

//go:embed stub_vpcmpeqw_128.go
var stubVpcmpeqw128 string

type VPCMPEQW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQW128() *VPCMPEQW128 {
	return &VPCMPEQW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 16),
	}
}

func (v *VPCMPEQW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQW128) Name() string {
	return "VPCMPEQW (128 bit) "
}

func (v *VPCMPEQW128) Description() string {
	return "Compare packed 16-bit integers for equality; result words are 0xFFFF if equal, else 0x0000."
}

func (v *VPCMPEQW128) Stub() string {
	return stubVpcmpeqw128
}

func (v *VPCMPEQW128) Assembly() string {
	return assemblyVpcmpeqw128
}

func (v *VPCMPEQW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpcmpeqw128(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
