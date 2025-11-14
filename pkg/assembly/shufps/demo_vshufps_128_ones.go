package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_ones.s
var assemblyVshufps128Ones string

//go:embed stub_vshufps_128_ones.go
var stubVshufps128Ones string

type VSHUFPS128ONES struct {
}

func (v *VSHUFPS128ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128ONES) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128ONES) Name() string {
	return "VSHUFPS (128 bit) ones"
}

func (v *VSHUFPS128ONES) Description() string {
	return "VSHUFPS imm8=0x55: dst = [a1,a1,b1,b1]"
}

func (v *VSHUFPS128ONES) Stub() string {
	return stubVshufps128Ones
}

func (v *VSHUFPS128ONES) Assembly() string {
	return assemblyVshufps128Ones
}

func (v *VSHUFPS128ONES) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Ones(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128ONES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
