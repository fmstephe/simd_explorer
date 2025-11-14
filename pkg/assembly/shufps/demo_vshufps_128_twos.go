package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_twos.s
var assemblyVshufps128Twos string

//go:embed stub_vshufps_128_twos.go
var stubVshufps128Twos string

type VSHUFPS128TWOS struct {
}

func (v *VSHUFPS128TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128TWOS) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128TWOS) Name() string {
	return "VSHUFPS (128 bit) twos"
}

func (v *VSHUFPS128TWOS) Description() string {
	return "VSHUFPS imm8=0xAA: dst = [a2,a2,b2,b2]"
}

func (v *VSHUFPS128TWOS) Stub() string {
	return stubVshufps128Twos
}

func (v *VSHUFPS128TWOS) Assembly() string {
	return assemblyVshufps128Twos
}

func (v *VSHUFPS128TWOS) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Twos(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128TWOS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
