package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_neq.s
var assemblyCmpps128Neq string

//go:embed stub_cmpps_128_neq.go
var stubCmpps128Neq string

type CMPPS128NEQ struct {
}

func (v *CMPPS128NEQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128NEQ) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128NEQ) Name() string {
	return "CMPPS (128 bit) neq"
}

func (v *CMPPS128NEQ) Description() string {
	return "Compare packed single-precision floats for not-equal (per lane)."
}

func (v *CMPPS128NEQ) Stub() string {
	return stubCmpps128Neq
}

func (v *CMPPS128NEQ) Assembly() string {
	return assemblyCmpps128Neq
}

func (v *CMPPS128NEQ) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Neq(&floats1, &floats2, &ret)

	log.Printf("CMPPS128NEQ input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128NEQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
