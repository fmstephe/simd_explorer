package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_andnps_128.s
var assemblyAndnps128 string

//go:embed stub_andnps_128.go
var stubAndnps128 string

type ANDNPS128 struct {
}

func (v *ANDNPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *ANDNPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *ANDNPS128) Name() string {
	return "ANDNPS (128 bit)"
}

func (v *ANDNPS128) Description() string {
	return "Bitwise AND NOT of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ANDNPS128) Stub() string {
	return stubAndnps128
}

func (v *ANDNPS128) Assembly() string {
	return assemblyAndnps128
}

func (v *ANDNPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	andnps128(&vals1, &vals2, &ret)

	log.Printf("ANDNPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *ANDNPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
