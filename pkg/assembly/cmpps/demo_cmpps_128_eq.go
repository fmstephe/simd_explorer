package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_eq.s
var assemblyCmpps128Eq string

//go:embed stub_cmpps_128_eq.go
var stubCmpps128Eq string

type CMPPS128EQ struct {
}

func (v *CMPPS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128EQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128EQ) Name() string {
	return "CMPPS (128 bit) eq"
}

func (v *CMPPS128EQ) Description() string {
	return "Compare packed single-precision floats for equality (per lane)."
}

func (v *CMPPS128EQ) Stub() string {
	return stubCmpps128Eq
}

func (v *CMPPS128EQ) Assembly() string {
	return assemblyCmpps128Eq
}

func (v *CMPPS128EQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Eq(&floats1, &floats2, &ret)

	log.Printf("CMPPS128EQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
