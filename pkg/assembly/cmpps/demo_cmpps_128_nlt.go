package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_nlt.s
var assemblyCmpps128Nlt string

//go:embed stub_cmpps_128_nlt.go
var stubCmpps128Nlt string

type CMPPS128NLT struct {
}

func (v *CMPPS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128NLT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128NLT) Name() string {
	return "CMPPS (128 bit) nlt"
}

func (v *CMPPS128NLT) Description() string {
	return "Compare packed single-precision floats for not-less-than (per lane)."
}

func (v *CMPPS128NLT) Stub() string {
	return stubCmpps128Nlt
}

func (v *CMPPS128NLT) Assembly() string {
	return assemblyCmpps128Nlt
}

func (v *CMPPS128NLT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Nlt(&floats1, &floats2, &ret)

	log.Printf("CMPPS128NLT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
