package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklqdq_256.s
var assemblyVpunpcklqdq256 string

//go:embed stub_vpunpcklqdq_256.go
var stubVpunpcklqdq256 string

type VPUNPCKLQDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLQDQ256() *VPUNPCKLQDQ256 {
	return &VPUNPCKLQDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPUNPCKLQDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLQDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLQDQ256) Name() string {
	return "VPUNPCKLQDQ (256 bit) "
}

func (v *VPUNPCKLQDQ256) Description() string {
	return "Unpack and interleave low-order quadwords from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKLQDQ256) Stub() string {
	return stubVpunpcklqdq256
}

func (v *VPUNPCKLQDQ256) Assembly() string {
	return assemblyVpunpcklqdq256
}

func (v *VPUNPCKLQDQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpunpcklqdq256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLQDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLQDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
