package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandnps_128.s
var assemblyVandnps128 string

//go:embed stub_vandnps_128.go
var stubVandnps128 string

type VANDNPS128 struct {
}

func (v *VANDNPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VANDNPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VANDNPS128) Name() string {
	return "VANDNPS (128 bit)"
}

func (v *VANDNPS128) Description() string {
	return "Bitwise AND NOT with VEX encoding; output shown as 32-bit hex lanes."
}

func (v *VANDNPS128) Stub() string {
	return stubVandnps128
}

func (v *VANDNPS128) Assembly() string {
	return assemblyVandnps128
}

func (v *VANDNPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vandnps128(&vals1, &vals2, &ret)

	log.Printf("VANDNPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VANDNPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
