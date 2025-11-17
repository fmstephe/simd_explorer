package andps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_andps_128.s
var assemblyAndps128 string

//go:embed stub_andps_128.go
var stubAndps128 string

type ANDPS128 struct {
}

func (v *ANDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *ANDPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *ANDPS128) Name() string {
	return "ANDPS (128 bit)"
}

func (v *ANDPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ANDPS128) Stub() string {
	return stubAndps128
}

func (v *ANDPS128) Assembly() string {
	return assemblyAndps128
}

func (v *ANDPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	andps128(&vals1, &vals2, &ret)

	log.Printf("ANDPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *ANDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
