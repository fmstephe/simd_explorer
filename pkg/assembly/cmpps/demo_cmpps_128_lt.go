package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_lt.s
var assemblyCmpps128Lt string

//go:embed stub_cmpps_128_lt.go
var stubCmpps128Lt string

type CMPPS128LT struct {
}

func (v *CMPPS128LT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128LT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128LT) Name() string {
	return "CMPPS (128 bit) lt"
}

func (v *CMPPS128LT) Description() string {
	return "Compare packed single-precision floats for less-than (per lane)."
}

func (v *CMPPS128LT) Stub() string {
	return stubCmpps128Lt
}

func (v *CMPPS128LT) Assembly() string {
	return assemblyCmpps128Lt
}

func (v *CMPPS128LT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Lt(&floats1, &floats2, &ret)

	log.Printf("CMPPS128LT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128LT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
