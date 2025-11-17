package andps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandps_128.s
var assemblyVandps128 string

//go:embed stub_vandps_128.go
var stubVandps128 string

type VANDPS128 struct {
}

func (v *VANDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VANDPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VANDPS128) Name() string {
	return "VANDPS (128 bit)"
}

func (v *VANDPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VANDPS128) Stub() string {
	return stubVandps128
}

func (v *VANDPS128) Assembly() string {
	return assemblyVandps128
}

func (v *VANDPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vandps128(&vals1, &vals2, &ret)

	log.Printf("VANDPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VANDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
