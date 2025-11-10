package comiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcomiss_128.s
var assemblyVcomiss128 string

//go:embed stub_vcomiss_128.go
var stubVcomiss128 string

type VCOMISS128 struct {
}

func (v *VCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCOMISS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VCOMISS128) Name() string {
	return "VCOMISS (128 bit)"
}

func (v *VCOMISS128) Description() string {
	return "TODO"
}

func (v *VCOMISS128) Stub() string {
	return stubVcomiss128
}

func (v *VCOMISS128) Assembly() string {
	return assemblyVcomiss128
}

func (v *VCOMISS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	flags := vcomiss128(&floats1, &floats2)

	log.Printf("VCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		floats1, floats2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	return []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
}

func (v *VCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
