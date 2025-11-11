package ucomiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_ucomiss_128.s
var assemblyComiss128 string

//go:embed stub_ucomiss_128.go
var stubComiss128 string

type UCOMISS128 struct {
}

func (v *UCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *UCOMISS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *UCOMISS128) Name() string {
	return "UCOMISS (128 bit)"
}

func (v *UCOMISS128) Description() string {
	return "TODO"
}

func (v *UCOMISS128) Stub() string {
	return stubComiss128
}

func (v *UCOMISS128) Assembly() string {
	return assemblyComiss128
}

func (v *UCOMISS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	flags := ucomiss128(&floats1, &floats2)

	log.Printf("UCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		floats1, floats2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	return []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
}

func (v *UCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
