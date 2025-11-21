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
}

func (v *VTESTPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VTESTPS256) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 16) // bit0=ZF, bit1=CF
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

func (v *VTESTPS256) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	var flags uint32

	vtestps256(&a, &b, &flags)

	log.Printf("VTESTPS256 A %v B %v flags 0x%X", a, b, flags)

	return number.Uint32ToBytes(flags)
}

func (v *VTESTPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
