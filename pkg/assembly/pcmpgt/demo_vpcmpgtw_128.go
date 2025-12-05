package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtw_128.s
var assemblyVpcmpgtw128 string

//go:embed stub_vpcmpgtw_128.go
var stubVpcmpgtw128 string

type VPCMPGTW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTW128() *VPCMPGTW128 {
	return &VPCMPGTW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 16),
	}
}

func (v *VPCMPGTW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTW128) Name() string {
	return "VPCMPGTW (128 bit) "
}

func (v *VPCMPGTW128) Description() string {
	return "Compare packed signed 16-bit integers for greater-than; words are 0xFFFF if vals1 > vals2 else 0x0000."
}

func (v *VPCMPGTW128) Stub() string {
	return stubVpcmpgtw128
}

func (v *VPCMPGTW128) Assembly() string {
	return assemblyVpcmpgtw128
}

func (v *VPCMPGTW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpcmpgtw128(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
