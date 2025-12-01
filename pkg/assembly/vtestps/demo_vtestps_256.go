package vtestps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vtestps_256.s
var assemblyVtestps256 string

//go:embed stub_vtestps_256.go
var stubVtestps256 string

type VTESTPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVTESTPS256() *VTESTPS256 {
	return &VTESTPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VTESTPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VTESTPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VTESTPS256) Name() string {
	return "VTESTPS (256 bit) "
}

func (v *VTESTPS256) Description() string {
	return "Test packed single-precision sign bits (per 128-bit lane). ZF=1 if (a & b) sign bits all zero; CF=1 if (~a & b) sign bits all zero. Output: bit0=ZF, bit1=CF."
}

func (v *VTESTPS256) Stub() string {
	return stubVtestps256
}

func (v *VTESTPS256) Assembly() string {
	return assemblyVtestps256
}

func (v *VTESTPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	var flags uint32
	vtestps256(&vals1, &vals2, &flags)

	log.Printf("VTESTPS256 vals1 %v vals2 %v flags 0x%X", vals1, vals2, flags)

	out := number.Uint32ToBytes(flags)
	v.ret.SetData(out)

}

func (v *VTESTPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
