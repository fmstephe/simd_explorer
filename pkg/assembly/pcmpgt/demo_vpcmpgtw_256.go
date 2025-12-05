package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtw_256.s
var assemblyVpcmpgtw256 string

//go:embed stub_vpcmpgtw_256.go
var stubVpcmpgtw256 string

type VPCMPGTW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTW256() *VPCMPGTW256 {
	return &VPCMPGTW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 16),
	}
}

func (v *VPCMPGTW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTW256) Name() string {
	return "VPCMPGTW (256 bit) "
}

func (v *VPCMPGTW256) Description() string {
	return "Compare packed signed 16-bit integers for greater-than; words are 0xFFFF if vals1 > vals2 else 0x0000."
}

func (v *VPCMPGTW256) Stub() string {
	return stubVpcmpgtw256
}

func (v *VPCMPGTW256) Assembly() string {
	return assemblyVpcmpgtw256
}

func (v *VPCMPGTW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpcmpgtw256(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPCMPGTW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
