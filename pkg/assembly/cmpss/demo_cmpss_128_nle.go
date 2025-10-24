package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_nle.s
var assemblyCmpss128Nle string

//go:embed stub_cmpss_128_nle.go
var stubCmpss128Nle string

type CMPSS128NLE struct {
}

func (v *CMPSS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128NLE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128NLE) Name() string {
	return "CMPSS (128 bit) nle"
}

func (v *CMPSS128NLE) Description() string {
	return "TODO"
}

func (v *CMPSS128NLE) Stub() string {
	return stubCmpss128Nle
}

func (v *CMPSS128NLE) Assembly() string {
	return assemblyCmpss128Nle
}

func (v *CMPSS128NLE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Nle(&floats1, &floats2, &ret)

	log.Printf("CMPSS128NLE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
