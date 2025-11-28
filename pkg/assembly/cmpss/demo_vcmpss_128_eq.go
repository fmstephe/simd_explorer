package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_eq.s
var assemblyVcmpss128Eq string

//go:embed stub_vcmpss_128_eq.go
var stubVcmpss128Eq string

type VCMPSS128EQ struct {
}

func (v *VCMPSS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128EQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128EQ) Name() string {
	return "VCMPSS (128 bit) eq"
}

func (v *VCMPSS128EQ) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for equality; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128EQ) Stub() string {
	return stubVcmpss128Eq
}

func (v *VCMPSS128EQ) Assembly() string {
	return assemblyVcmpss128Eq
}

func (v *VCMPSS128EQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Eq(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128EQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
