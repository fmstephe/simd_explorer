package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_unord.s
var assemblyCmpps128Unord string

//go:embed stub_cmpps_128_unord.go
var stubCmpps128Unord string

type CMPPS128UNORD struct {
}

func (v *CMPPS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *CMPPS128UNORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *CMPPS128UNORD) Name() string {
	return "CMPPS (128 bit) unord"
}

func (v *CMPPS128UNORD) Description() string {
	return "Compare packed single-precision floats for unordered (per lane)."
}

func (v *CMPPS128UNORD) Stub() string {
	return stubCmpps128Unord
}

func (v *CMPPS128UNORD) Assembly() string {
	return assemblyCmpps128Unord
}

func (v *CMPPS128UNORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	cmpps128Unord(&floats1, &floats2, &ret)

	log.Printf("CMPPS128UNORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *CMPPS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
