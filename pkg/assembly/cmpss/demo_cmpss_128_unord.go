package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_unord.s
var assemblyCmpss128Unord string

//go:embed stub_cmpss_128_unord.go
var stubCmpss128Unord string

type CMPSS128UNORD struct {
}

func (v *CMPSS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128UNORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128UNORD) Name() string {
	return "CMPSS (128 bit) unord"
}

func (v *CMPSS128UNORD) Description() string {
	return "Compare scalar single-precision (lane 0) for unordered (either operand is NaN); result mask in lane 0."
}

func (v *CMPSS128UNORD) Stub() string {
	return stubCmpss128Unord
}

func (v *CMPSS128UNORD) Assembly() string {
	return assemblyCmpss128Unord
}

func (v *CMPSS128UNORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Unord(&floats1, &floats2, &ret)

	log.Printf("CMPSS128UNORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
