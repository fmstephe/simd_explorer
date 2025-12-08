package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhqdq_256.s
var assemblyVpunpckhqdq256 string

//go:embed stub_vpunpckhqdq_256.go
var stubVpunpckhqdq256 string

type VPUNPCKHQDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHQDQ256() *VPUNPCKHQDQ256 {
	return &VPUNPCKHQDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPUNPCKHQDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHQDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHQDQ256) Name() string {
	return "VPUNPCKHQDQ (256 bit) "
}

func (v *VPUNPCKHQDQ256) Description() string {
	return "Unpack and interleave high-order quadwords from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKHQDQ256) Stub() string {
	return stubVpunpckhqdq256
}

func (v *VPUNPCKHQDQ256) Assembly() string {
	return assemblyVpunpckhqdq256
}

func (v *VPUNPCKHQDQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpunpckhqdq256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHQDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHQDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
