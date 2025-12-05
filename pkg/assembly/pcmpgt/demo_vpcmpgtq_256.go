package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtq_256.s
var assemblyVpcmpgtq256 string

//go:embed stub_vpcmpgtq_256.go
var stubVpcmpgtq256 string

type VPCMPGTQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTQ256() *VPCMPGTQ256 {
	return &VPCMPGTQ256{
		vals1: number.NewNamedIntParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPCMPGTQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTQ256) Name() string {
	return "VPCMPGTQ (256 bit) "
}

func (v *VPCMPGTQ256) Description() string {
	return "Compare packed signed 64-bit integers for greater-than; qwords are all-ones if vals1 > vals2 else zero."
}

func (v *VPCMPGTQ256) Stub() string {
	return stubVpcmpgtq256
}

func (v *VPCMPGTQ256) Assembly() string {
	return assemblyVpcmpgtq256
}

func (v *VPCMPGTQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpcmpgtq256(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
