package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtq_128.s
var assemblyVpcmpgtq128 string

//go:embed stub_vpcmpgtq_128.go
var stubVpcmpgtq128 string

type VPCMPGTQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTQ128() *VPCMPGTQ128 {
	return &VPCMPGTQ128{
		vals1: number.NewNamedIntParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 16),
	}
}

func (v *VPCMPGTQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTQ128) Name() string {
	return "VPCMPGTQ (128 bit) "
}

func (v *VPCMPGTQ128) Description() string {
	return "Compare packed signed 64-bit integers for greater-than; qwords are all-ones if vals1 > vals2 else zero."
}

func (v *VPCMPGTQ128) Stub() string {
	return stubVpcmpgtq128
}

func (v *VPCMPGTQ128) Assembly() string {
	return assemblyVpcmpgtq128
}

func (v *VPCMPGTQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpcmpgtq128(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
