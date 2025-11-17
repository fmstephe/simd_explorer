package andps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandps_256.s
var assemblyVandps256 string

//go:embed stub_vandps_256.go
var stubVandps256 string

type VANDPS256 struct {
}

func (v *VANDPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VANDPS256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VANDPS256) Name() string {
	return "VANDPS (256 bit)"
}

func (v *VANDPS256) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VANDPS256) Stub() string {
	return stubVandps256
}

func (v *VANDPS256) Assembly() string {
	return assemblyVandps256
}

func (v *VANDPS256) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vandps256(&vals1, &vals2, &ret)

	log.Printf("VANDPS256 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VANDPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
