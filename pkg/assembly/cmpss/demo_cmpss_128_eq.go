package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_eq.s
var assemblyCmpss128Eq string

//go:embed stub_cmpss_128_eq.go
var stubCmpss128Eq string

type CMPSS128EQ struct {
}

func (v *CMPSS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128EQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128EQ) Name() string {
	return "CMPSS (128 bit) eq"
}

func (v *CMPSS128EQ) Description() string {
	return "Compare scalar single-precision (lane 0) for equality; writes 0xFFFFFFFF or 0x0 in lane 0, upper lanes pass through."
}

func (v *CMPSS128EQ) Stub() string {
	return stubCmpss128Eq
}

func (v *CMPSS128EQ) Assembly() string {
	return assemblyCmpss128Eq
}

func (v *CMPSS128EQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Eq(&floats1, &floats2, &ret)

	log.Printf("CMPSS128EQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
