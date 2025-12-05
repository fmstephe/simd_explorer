package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqd_256.s
var assemblyVpcmpeqd256 string

//go:embed stub_vpcmpeqd_256.go
var stubVpcmpeqd256 string

type VPCMPEQD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQD256() *VPCMPEQD256 {
	return &VPCMPEQD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VPCMPEQD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQD256) Name() string {
	return "VPCMPEQD (256 bit) "
}

func (v *VPCMPEQD256) Description() string {
	return "Compare packed 32-bit integers for equality; result dwords are 0xFFFFFFFF if equal, else 0x00000000."
}

func (v *VPCMPEQD256) Stub() string {
	return stubVpcmpeqd256
}

func (v *VPCMPEQD256) Assembly() string {
	return assemblyVpcmpeqd256
}

func (v *VPCMPEQD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpcmpeqd256(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
