package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_lt.s
var assemblyCmpss128Lt string

//go:embed stub_cmpss_128_lt.go
var stubCmpss128Lt string

type CMPSS128LT struct {
}

func (v *CMPSS128LT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128LT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128LT) Name() string {
	return "CMPSS (128 bit) lt"
}

func (v *CMPSS128LT) Description() string {
	return "Compare scalar single-precision (lane 0) for less-than; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128LT) Stub() string {
	return stubCmpss128Lt
}

func (v *CMPSS128LT) Assembly() string {
	return assemblyCmpss128Lt
}

func (v *CMPSS128LT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Lt(&floats1, &floats2, &ret)

	log.Printf("CMPSS128LT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128LT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
