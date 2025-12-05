package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqd_128.s
var assemblyVpcmpeqd128 string

//go:embed stub_vpcmpeqd_128.go
var stubVpcmpeqd128 string

type VPCMPEQD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQD128() *VPCMPEQD128 {
	return &VPCMPEQD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPCMPEQD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQD128) Name() string {
	return "VPCMPEQD (128 bit) "
}

func (v *VPCMPEQD128) Description() string {
	return "Compare packed 32-bit integers for equality; result dwords are 0xFFFFFFFF if equal, else 0x00000000."
}

func (v *VPCMPEQD128) Stub() string {
	return stubVpcmpeqd128
}

func (v *VPCMPEQD128) Assembly() string {
	return assemblyVpcmpeqd128
}

func (v *VPCMPEQD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpcmpeqd128(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
