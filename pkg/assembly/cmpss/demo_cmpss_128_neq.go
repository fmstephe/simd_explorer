package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_neq.s
var assemblyCmpss128Neq string

//go:embed stub_cmpss_128_neq.go
var stubCmpss128Neq string

type CMPSS128NEQ struct {
}

func (v *CMPSS128NEQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128NEQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128NEQ) Name() string {
	return "CMPSS (128 bit) neq"
}

func (v *CMPSS128NEQ) Description() string {
	return "TODO"
}

func (v *CMPSS128NEQ) Stub() string {
	return stubCmpss128Neq
}

func (v *CMPSS128NEQ) Assembly() string {
	return assemblyCmpss128Neq
}

func (v *CMPSS128NEQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Neq(&floats1, &floats2, &ret)

	log.Printf("CMPSS128NEQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128NEQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
