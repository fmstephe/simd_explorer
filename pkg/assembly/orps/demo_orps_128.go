package orps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_orps_128.s
var assemblyOrps128 string

//go:embed stub_orps_128.go
var stubOrps128 string

type ORPS128 struct {
}

func (v *ORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *ORPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *ORPS128) Name() string {
	return "ORPS (128 bit)"
}

func (v *ORPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ORPS128) Stub() string {
	return stubOrps128
}

func (v *ORPS128) Assembly() string {
	return assemblyOrps128
}

func (v *ORPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	orps128(&vals1, &vals2, &ret)

	log.Printf("ORPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *ORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
