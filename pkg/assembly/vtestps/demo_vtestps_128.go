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
}

func (v *VTESTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VTESTPS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 16) // bit0=ZF, bit1=CF (hex)
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

func (v *VTESTPS128) Run(inputs [][]byte) (output []byte) {
	a := [4]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [4]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	var flags uint32

	vtestps128(&a, &b, &flags)

	log.Printf("VTESTPS128 A %v B %v flags 0x%X", a, b, flags)

	return number.Uint32ToBytes(flags)
}

func (v *VTESTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
