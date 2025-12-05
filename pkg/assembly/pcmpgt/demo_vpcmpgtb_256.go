package pcmpgt

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpcmpgtb_256.s
var assemblyVpcmpgtb256 string

//go:embed stub_vpcmpgtb_256.go
var stubVpcmpgtb256 string

type VPCMPGTB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPCMPGTB256() *VPCMPGTB256 {
	return &VPCMPGTB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 16),
	}
}

func (v *VPCMPGTB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPCMPGTB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPCMPGTB256) Name() string {
	return "VPCMPGTB (256 bit) "
}

func (v *VPCMPGTB256) Description() string {
	return "Compare packed signed 8-bit integers for greater-than; bytes are 0xFF if vals1 > vals2 else 0x00."
}

func (v *VPCMPGTB256) Stub() string {
	return stubVpcmpgtb256
}

func (v *VPCMPGTB256) Assembly() string {
	return assemblyVpcmpgtb256
}

func (v *VPCMPGTB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpcmpgtb256(&vals1, &vals2, &ret)

	log.Printf("VPCMPGTB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPCMPGTB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
