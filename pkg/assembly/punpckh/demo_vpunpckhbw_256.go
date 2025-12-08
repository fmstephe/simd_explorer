package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhbw_256.s
var assemblyVpunpckhbw256 string

//go:embed stub_vpunpckhbw_256.go
var stubVpunpckhbw256 string

type VPUNPCKHBW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHBW256() *VPUNPCKHBW256 {
	return &VPUNPCKHBW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPUNPCKHBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHBW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHBW256) Name() string {
	return "VPUNPCKHBW (256 bit) "
}

func (v *VPUNPCKHBW256) Description() string {
	return "Unpack and interleave high-order bytes from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKHBW256) Stub() string {
	return stubVpunpckhbw256
}

func (v *VPUNPCKHBW256) Assembly() string {
	return assemblyVpunpckhbw256
}

func (v *VPUNPCKHBW256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpunpckhbw256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHBW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPUNPCKHBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
