package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklbw_256.s
var assemblyVpunpcklbw256 string

//go:embed stub_vpunpcklbw_256.go
var stubVpunpcklbw256 string

type VPUNPCKLBW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLBW256() *VPUNPCKLBW256 {
	return &VPUNPCKLBW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPUNPCKLBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		// TODO replace with actual parameters for instruction demo
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLBW256) Output() *number.Parameter {
	// TODO replace with actual parameters for instruction demo
	return v.ret
}

func (v *VPUNPCKLBW256) Name() string {
	return "VPUNPCKLBW (256 bit) "
}

func (v *VPUNPCKLBW256) Description() string {
	return "Unpack and interleave low-order bytes from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKLBW256) Stub() string {
	return stubVpunpcklbw256
}

func (v *VPUNPCKLBW256) Assembly() string {
	return assemblyVpunpcklbw256
}

func (v *VPUNPCKLBW256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpunpcklbw256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLBW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPUNPCKLBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
