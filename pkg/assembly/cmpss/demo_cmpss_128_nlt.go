package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_nlt.s
var assemblyCmpss128Nlt string

//go:embed stub_cmpss_128_nlt.go
var stubCmpss128Nlt string

type CMPSS128NLT struct {
}

func (v *CMPSS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128NLT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128NLT) Name() string {
	return "CMPSS (128 bit) nlt"
}

func (v *CMPSS128NLT) Description() string {
	return "Compare scalar single-precision (lane 0) for not-less-than; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128NLT) Stub() string {
	return stubCmpss128Nlt
}

func (v *CMPSS128NLT) Assembly() string {
	return assemblyCmpss128Nlt
}

func (v *CMPSS128NLT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Nlt(&floats1, &floats2, &ret)

	log.Printf("CMPSS128NLT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
