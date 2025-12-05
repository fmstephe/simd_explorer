package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtd_128.s
var assemblyVpcmpgtd128 string

//go:embed stub_vpcmpgtd_128.go
var stubVpcmpgtd128 string

type VPCMPGTD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTD128() *VPCMPGTD128 {
	return &VPCMPGTD128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPCMPGTD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTD128) Name() string {
	return "VPCMPGTD (128 bit) "
}

func (v *VPCMPGTD128) Description() string {
	return "Compare packed signed 32-bit integers for greater-than; dwords are 0xFFFFFFFF if vals1 > vals2 else 0x00000000."
}

func (v *VPCMPGTD128) Stub() string {
	return stubVpcmpgtd128
}

func (v *VPCMPGTD128) Assembly() string {
	return assemblyVpcmpgtd128
}

func (v *VPCMPGTD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpcmpgtd128(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
