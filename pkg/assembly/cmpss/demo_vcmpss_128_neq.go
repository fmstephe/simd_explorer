package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_neq.s
var assemblyVcmpss128Neq string

//go:embed stub_vcmpss_128_neq.go
var stubVcmpss128Neq string

type VCMPSS128NEQ struct {
}

func (v *VCMPSS128NEQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128NEQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128NEQ) Name() string {
	return "VCMPSS (128 bit) neq"
}

func (v *VCMPSS128NEQ) Description() string {
	return "TODO"
}

func (v *VCMPSS128NEQ) Stub() string {
	return stubVcmpss128Neq
}

func (v *VCMPSS128NEQ) Assembly() string {
	return assemblyVcmpss128Neq
}

func (v *VCMPSS128NEQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Neq(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128NEQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128NEQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
