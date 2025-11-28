package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_le.s
var assemblyCmpss128Le string

//go:embed stub_cmpss_128_le.go
var stubCmpss128Le string

type CMPSS128LE struct {
}

func (v *CMPSS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128LE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128LE) Name() string {
	return "CMPSS (128 bit) le"
}

func (v *CMPSS128LE) Description() string {
	return "Compare scalar single-precision (lane 0) for less-than-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128LE) Stub() string {
	return stubCmpss128Le
}

func (v *CMPSS128LE) Assembly() string {
	return assemblyCmpss128Le
}

func (v *CMPSS128LE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Le(&floats1, &floats2, &ret)

	log.Printf("CMPSS128LE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
