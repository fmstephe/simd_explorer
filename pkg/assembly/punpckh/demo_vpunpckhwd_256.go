package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhwd_256.s
var assemblyVpunpckhwd256 string

//go:embed stub_vpunpckhwd_256.go
var stubVpunpckhwd256 string

type VPUNPCKHWD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHWD256() *VPUNPCKHWD256 {
	return &VPUNPCKHWD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPUNPCKHWD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHWD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHWD256) Name() string {
	return "VPUNPCKHWD (256 bit) "
}

func (v *VPUNPCKHWD256) Description() string {
	return "Unpack and interleave high-order words from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKHWD256) Stub() string {
	return stubVpunpckhwd256
}

func (v *VPUNPCKHWD256) Assembly() string {
	return assemblyVpunpckhwd256
}

func (v *VPUNPCKHWD256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpunpckhwd256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHWD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHWD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
