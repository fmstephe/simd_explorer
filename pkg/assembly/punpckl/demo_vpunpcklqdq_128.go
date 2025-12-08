package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklqdq_128.s
var assemblyVpunpcklqdq128 string

//go:embed stub_vpunpcklqdq_128.go
var stubVpunpcklqdq128 string

type VPUNPCKLQDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLQDQ128() *VPUNPCKLQDQ128 {
	return &VPUNPCKLQDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPUNPCKLQDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLQDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLQDQ128) Name() string {
	return "VPUNPCKLQDQ (128 bit) "
}

func (v *VPUNPCKLQDQ128) Description() string {
	return "Unpack and interleave low-order quadwords from two 128-bit sources."
}

func (v *VPUNPCKLQDQ128) Stub() string {
	return stubVpunpcklqdq128
}

func (v *VPUNPCKLQDQ128) Assembly() string {
	return assemblyVpunpcklqdq128
}

func (v *VPUNPCKLQDQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpunpcklqdq128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLQDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLQDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
