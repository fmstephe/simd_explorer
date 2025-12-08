package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhbw_128.s
var assemblyVpunpckhbw128 string

//go:embed stub_vpunpckhbw_128.go
var stubVpunpckhbw128 string

type VPUNPCKHBW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHBW128() *VPUNPCKHBW128 {
	return &VPUNPCKHBW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPUNPCKHBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHBW128) Name() string {
	return "VPUNPCKHBW (128 bit) "
}

func (v *VPUNPCKHBW128) Description() string {
	return "Unpack and interleave high-order bytes from two 128-bit sources."
}

func (v *VPUNPCKHBW128) Stub() string {
	return stubVpunpckhbw128
}

func (v *VPUNPCKHBW128) Assembly() string {
	return assemblyVpunpckhbw128
}

func (v *VPUNPCKHBW128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpunpckhbw128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHBW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPUNPCKHBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
