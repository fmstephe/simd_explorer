package pcmpeq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpeqq_256.s
var assemblyVpcmpeqq256 string

//go:embed stub_vpcmpeqq_256.go
var stubVpcmpeqq256 string

type VPCMPEQQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPEQQ256() *VPCMPEQQ256 {
	return &VPCMPEQQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPCMPEQQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPEQQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPEQQ256) Name() string {
	return "VPCMPEQQ (256 bit) "
}

func (v *VPCMPEQQ256) Description() string {
	return "Compare packed 64-bit integers for equality; result qwords are all-ones if equal, else zero."
}

func (v *VPCMPEQQ256) Stub() string {
	return stubVpcmpeqq256
}

func (v *VPCMPEQQ256) Assembly() string {
	return assemblyVpcmpeqq256
}

func (v *VPCMPEQQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpcmpeqq256(&vals1, &vals2, &ret)

	log.Printf("VPCMPEQQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPEQQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
