package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtd_256.s
var assemblyVpcmpgtd256 string

//go:embed stub_vpcmpgtd_256.go
var stubVpcmpgtd256 string

type VPCMPGTD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTD256() *VPCMPGTD256 {
	return &VPCMPGTD256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VPCMPGTD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTD256) Name() string {
	return "VPCMPGTD (256 bit) "
}

func (v *VPCMPGTD256) Description() string {
	return "Compare packed signed 32-bit integers for greater-than; dwords are 0xFFFFFFFF if vals1 > vals2 else 0x00000000."
}

func (v *VPCMPGTD256) Stub() string {
	return stubVpcmpgtd256
}

func (v *VPCMPGTD256) Assembly() string {
	return assemblyVpcmpgtd256
}

func (v *VPCMPGTD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpcmpgtd256(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
