package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_unord.s
var assemblyVcmpss128Unord string

//go:embed stub_vcmpss_128_unord.go
var stubVcmpss128Unord string

type VCMPSS128UNORD struct {
}

func (v *VCMPSS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128UNORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VCMPSS128UNORD) Name() string {
	return "VCMPSS (128 bit) unord"
}

func (v *VCMPSS128UNORD) Description() string {
	return "TODO"
}

func (v *VCMPSS128UNORD) Stub() string {
	return stubVcmpss128Unord
}

func (v *VCMPSS128UNORD) Assembly() string {
	return assemblyVcmpss128Unord
}

func (v *VCMPSS128UNORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Unord(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128UNORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
