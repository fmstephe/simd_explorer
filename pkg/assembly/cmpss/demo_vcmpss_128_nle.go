package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_nle.s
var assemblyVcmpss128Nle string

//go:embed stub_vcmpss_128_nle.go
var stubVcmpss128Nle string

type VCMPSS128NLE struct {
}

func (v *VCMPSS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128NLE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128NLE) Name() string {
	return "VCMPSS (128 bit) nle"
}

func (v *VCMPSS128NLE) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for not-less-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128NLE) Stub() string {
	return stubVcmpss128Nle
}

func (v *VCMPSS128NLE) Assembly() string {
	return assemblyVcmpss128Nle
}

func (v *VCMPSS128NLE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Nle(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128NLE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
