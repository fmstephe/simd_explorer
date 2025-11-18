package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandnps_256.s
var assemblyVandnps256 string

//go:embed stub_vandnps_256.go
var stubVandnps256 string

type VANDNPS256 struct {
}

func (v *VANDNPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VANDNPS256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VANDNPS256) Name() string {
	return "VANDNPS (256 bit)"
}

func (v *VANDNPS256) Description() string {
	return "Bitwise AND NOT with VEX encoding (per 128-bit lane); output shown as 32-bit hex lanes."
}

func (v *VANDNPS256) Stub() string {
	return stubVandnps256
}

func (v *VANDNPS256) Assembly() string {
	return assemblyVandnps256
}

func (v *VANDNPS256) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vandnps256(&vals1, &vals2, &ret)

	log.Printf("VANDNPS256 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VANDNPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
