package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_nle.s
var assemblyCmpps128Nle string

//go:embed stub_cmpps_128_nle.go
var stubCmpps128Nle string

type CMPPS128NLE struct {
}

func (v *CMPPS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128NLE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128NLE) Name() string {
	return "CMPPS (128 bit) nle"
}

func (v *CMPPS128NLE) Description() string {
	return "Compare packed single-precision floats for not-less-or-equal (per lane)."
}

func (v *CMPPS128NLE) Stub() string {
	return stubCmpps128Nle
}

func (v *CMPPS128NLE) Assembly() string {
	return assemblyCmpps128Nle
}

func (v *CMPPS128NLE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Nle(&floats1, &floats2, &ret)

	log.Printf("CMPPS128NLE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
