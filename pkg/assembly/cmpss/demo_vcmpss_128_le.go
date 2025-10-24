package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_le.s
var assemblyVcmpss128Le string

//go:embed stub_vcmpss_128_le.go
var stubVcmpss128Le string

type VCMPSS128LE struct {
}

func (v *VCMPSS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128LE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128LE) Name() string {
	return "VCMPSS (128 bit) le"
}

func (v *VCMPSS128LE) Description() string {
	return "TODO"
}

func (v *VCMPSS128LE) Stub() string {
	return stubVcmpss128Le
}

func (v *VCMPSS128LE) Assembly() string {
	return assemblyVcmpss128Le
}

func (v *VCMPSS128LE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Le(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128LE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
