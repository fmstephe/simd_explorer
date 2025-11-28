package ucomiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vucomiss_128.s
var assemblyVucomiss128 string

//go:embed stub_vucomiss_128.go
var stubVucomiss128 string

type VUCOMISS128 struct {
}

func (v *VUCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VUCOMISS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VUCOMISS128) Name() string {
	return "VUCOMISS (128 bit)"
}

func (v *VUCOMISS128) Description() string {
	return "AVX form of unordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF)."
}

func (v *VUCOMISS128) Stub() string {
	return stubVucomiss128
}

func (v *VUCOMISS128) Assembly() string {
	return assemblyVucomiss128
}

func (v *VUCOMISS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	flags := vucomiss128(&floats1, &floats2)

	log.Printf("VUCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		floats1, floats2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	return []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
}

func (v *VUCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
