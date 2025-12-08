package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklwd_256.s
var assemblyVpunpcklwd256 string

//go:embed stub_vpunpcklwd_256.go
var stubVpunpcklwd256 string

type VPUNPCKLWD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLWD256() *VPUNPCKLWD256 {
	return &VPUNPCKLWD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPUNPCKLWD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLWD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLWD256) Name() string {
	return "VPUNPCKLWD (256 bit) "
}

func (v *VPUNPCKLWD256) Description() string {
	return "Unpack and interleave low-order words from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKLWD256) Stub() string {
	return stubVpunpcklwd256
}

func (v *VPUNPCKLWD256) Assembly() string {
	return assemblyVpunpcklwd256
}

func (v *VPUNPCKLWD256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpunpcklwd256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLWD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLWD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
