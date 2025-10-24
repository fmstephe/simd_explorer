package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_nlt.s
var assemblyVcmpss128Nlt string

//go:embed stub_vcmpss_128_nlt.go
var stubVcmpss128Nlt string

type VCMPSS128NLT struct {
}

func (v *VCMPSS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128NLT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128NLT) Name() string {
	return "VCMPSS (128 bit) nlt"
}

func (v *VCMPSS128NLT) Description() string {
	return "TODO"
}

func (v *VCMPSS128NLT) Stub() string {
	return stubVcmpss128Nlt
}

func (v *VCMPSS128NLT) Assembly() string {
	return assemblyVcmpss128Nlt
}

func (v *VCMPSS128NLT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Nlt(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128NLT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
