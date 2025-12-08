package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklbw_128.s
var assemblyVpunpcklbw128 string

//go:embed stub_vpunpcklbw_128.go
var stubVpunpcklbw128 string

type VPUNPCKLBW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLBW128() *VPUNPCKLBW128 {
	return &VPUNPCKLBW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPUNPCKLBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLBW128) Name() string {
	return "VPUNPCKLBW (128 bit) "
}

func (v *VPUNPCKLBW128) Description() string {
	return "Unpack and interleave low-order bytes from two 128-bit sources."
}

func (v *VPUNPCKLBW128) Stub() string {
	return stubVpunpcklbw128
}

func (v *VPUNPCKLBW128) Assembly() string {
	return assemblyVpunpcklbw128
}

func (v *VPUNPCKLBW128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpunpcklbw128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLBW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPUNPCKLBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
