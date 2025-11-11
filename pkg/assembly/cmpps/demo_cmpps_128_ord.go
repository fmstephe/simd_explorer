package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_ord.s
var assemblyCmpps128Ord string

//go:embed stub_cmpps_128_ord.go
var stubCmpps128Ord string

type CMPPS128ORD struct {
}

func (v *CMPPS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128ORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128ORD) Name() string {
	return "CMPPS (128 bit) ord"
}

func (v *CMPPS128ORD) Description() string {
	return "Compare packed single-precision floats for ordered (per lane)."
}

func (v *CMPPS128ORD) Stub() string {
	return stubCmpps128Ord
}

func (v *CMPPS128ORD) Assembly() string {
	return assemblyCmpps128Ord
}

func (v *CMPPS128ORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Ord(&floats1, &floats2, &ret)

	log.Printf("CMPPS128ORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
