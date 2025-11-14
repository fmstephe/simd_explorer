package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_ones.s
var assemblyVshufps256Ones string

//go:embed stub_vshufps_256_ones.go
var stubVshufps256Ones string

type VSHUFPS256ONES struct {
}

func (v *VSHUFPS256ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256ONES) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256ONES) Name() string {
	return "VSHUFPS (256 bit) ones"
}

func (v *VSHUFPS256ONES) Description() string {
	return "VSHUFPS imm8=0x55 per 128-bit lane: dst = [a1,a1,b1,b1 | a5,a5,b5,b5]"
}

func (v *VSHUFPS256ONES) Stub() string {
	return stubVshufps256Ones
}

func (v *VSHUFPS256ONES) Assembly() string {
	return assemblyVshufps256Ones
}

func (v *VSHUFPS256ONES) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Ones(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256ONES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
