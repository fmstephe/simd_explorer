package comiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_comiss_128.s
var assemblyComiss128 string

//go:embed stub_comiss_128.go
var stubComiss128 string

type COMISS128 struct {
}

func (v *COMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *COMISS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *COMISS128) Name() string {
	return "COMISS (128 bit)"
}

func (v *COMISS128) Description() string {
	return "Ordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF)."
}

func (v *COMISS128) Stub() string {
	return stubComiss128
}

func (v *COMISS128) Assembly() string {
	return assemblyComiss128
}

func (v *COMISS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	flags := comiss128(&floats1, &floats2)

	log.Printf("COMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		floats1, floats2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	return []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
}

func (v *COMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
