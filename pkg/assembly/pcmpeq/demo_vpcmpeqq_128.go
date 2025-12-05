package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqq_128.s
var assemblyVpcmpeqq128 string

//go:embed stub_vpcmpeqq_128.go
var stubVpcmpeqq128 string

type VPCMPEQQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQQ128() *VPCMPEQQ128 {
	return &VPCMPEQQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 16),
	}
}

func (v *VPCMPEQQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQQ128) Name() string {
	return "VPCMPEQQ (128 bit) "
}

func (v *VPCMPEQQ128) Description() string {
	return "Compare packed 64-bit integers for equality; result qwords are all-ones if equal, else zero."
}

func (v *VPCMPEQQ128) Stub() string {
	return stubVpcmpeqq128
}

func (v *VPCMPEQQ128) Assembly() string {
	return assemblyVpcmpeqq128
}

func (v *VPCMPEQQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpcmpeqq128(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
