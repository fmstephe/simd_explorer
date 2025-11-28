package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_lt.s
var assemblyVcmpss128Lt string

//go:embed stub_vcmpss_128_lt.go
var stubVcmpss128Lt string

type VCMPSS128LT struct {
}

func (v *VCMPSS128LT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128LT) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128LT) Name() string {
	return "VCMPSS (128 bit) lt"
}

func (v *VCMPSS128LT) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for less-than; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128LT) Stub() string {
	return stubVcmpss128Lt
}

func (v *VCMPSS128LT) Assembly() string {
	return assemblyVcmpss128Lt
}

func (v *VCMPSS128LT) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Lt(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128LT input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128LT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
