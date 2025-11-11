package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_le.s
var assemblyCmpps128Le string

//go:embed stub_cmpps_128_le.go
var stubCmpps128Le string

type CMPPS128LE struct {
}

func (v *CMPPS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128LE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128LE) Name() string {
	return "CMPPS (128 bit) le"
}

func (v *CMPPS128LE) Description() string {
	return "Compare packed single-precision floats for less-than-or-equal (per lane)."
}

func (v *CMPPS128LE) Stub() string {
	return stubCmpps128Le
}

func (v *CMPPS128LE) Assembly() string {
	return assemblyCmpps128Le
}

func (v *CMPPS128LE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Le(&floats1, &floats2, &ret)

	log.Printf("CMPPS128LE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
