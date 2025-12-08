package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhqdq_128.s
var assemblyVpunpckhqdq128 string

//go:embed stub_vpunpckhqdq_128.go
var stubVpunpckhqdq128 string

type VPUNPCKHQDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHQDQ128() *VPUNPCKHQDQ128 {
	return &VPUNPCKHQDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPUNPCKHQDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHQDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHQDQ128) Name() string {
	return "VPUNPCKHQDQ (128 bit) "
}

func (v *VPUNPCKHQDQ128) Description() string {
	return "Unpack and interleave high-order quadwords from two 128-bit sources."
}

func (v *VPUNPCKHQDQ128) Stub() string {
	return stubVpunpckhqdq128
}

func (v *VPUNPCKHQDQ128) Assembly() string {
	return assemblyVpunpckhqdq128
}

func (v *VPUNPCKHQDQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpunpckhqdq128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHQDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHQDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
