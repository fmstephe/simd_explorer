package orps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vorps_256.s
var assemblyVorps256 string

//go:embed stub_vorps_256.go
var stubVorps256 string

type VORPS256 struct {
}

func (v *VORPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VORPS256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VORPS256) Name() string {
	return "VORPS (256 bit)"
}

func (v *VORPS256) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VORPS256) Stub() string {
	return stubVorps256
}

func (v *VORPS256) Assembly() string {
	return assemblyVorps256
}

func (v *VORPS256) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vorps256(&vals1, &vals2, &ret)

	log.Printf("VORPS256 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VORPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
