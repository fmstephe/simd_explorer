package vtestps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vtestps_128.s
var assemblyVtestps128 string

//go:embed stub_vtestps_128.go
var stubVtestps128 string

type VTESTPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVTESTPS128() *VTESTPS128 {
	return &VTESTPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VTESTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals1, v.vals2}
}

func (v *VTESTPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VTESTPS128) Name() string {
	return "VTESTPS (128 bit) "
}

func (v *VTESTPS128) Description() string {
	return "Test packed single-precision sign bits: sets ZF if (a & b) sign bits all zero; CF if (~a & b) sign bits all zero. Output: bit0=ZF, bit1=CF."
}

func (v *VTESTPS128) Stub() string {
	return stubVtestps128
}

func (v *VTESTPS128) Assembly() string {
	return assemblyVtestps128
}

func (v *VTESTPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	var flags uint32
	vtestps128(&vals1, &vals2, &flags)

	log.Printf("VTESTPS128 vals1 %v vals2 %v flags 0x%X", vals1, vals2, flags)

	out := number.Uint32ToBytes(flags)
	v.ret.SetData(out)
	return out
}

func (v *VTESTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
