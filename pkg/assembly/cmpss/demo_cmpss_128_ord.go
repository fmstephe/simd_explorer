package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_ord.s
var assemblyCmpss128Ord string

//go:embed stub_cmpss_128_ord.go
var stubCmpss128Ord string

type CMPSS128ORD struct {
}

func (v *CMPSS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPSS128ORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPSS128ORD) Name() string {
	return "CMPSS (128 bit) ord"
}

func (v *CMPSS128ORD) Description() string {
	return "TODO"
}

func (v *CMPSS128ORD) Stub() string {
	return stubCmpss128Ord
}

func (v *CMPSS128ORD) Assembly() string {
	return assemblyCmpss128Ord
}

func (v *CMPSS128ORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpss128Ord(&floats1, &floats2, &ret)

	log.Printf("CMPSS128ORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPSS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
