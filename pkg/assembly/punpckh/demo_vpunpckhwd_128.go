package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhwd_128.s
var assemblyVpunpckhwd128 string

//go:embed stub_vpunpckhwd_128.go
var stubVpunpckhwd128 string

type VPUNPCKHWD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHWD128() *VPUNPCKHWD128 {
	return &VPUNPCKHWD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPUNPCKHWD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHWD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHWD128) Name() string {
	return "VPUNPCKHWD (128 bit) "
}

func (v *VPUNPCKHWD128) Description() string {
	return "Unpack and interleave high-order words from two 128-bit sources."
}

func (v *VPUNPCKHWD128) Stub() string {
	return stubVpunpckhwd128
}

func (v *VPUNPCKHWD128) Assembly() string {
	return assemblyVpunpckhwd128
}

func (v *VPUNPCKHWD128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpunpckhwd128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHWD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHWD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
