package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_sqrtps_128.s
var assemblySqrtps128 string

//go:embed stub_sqrtps_128.go
var stubSqrtps128 string

type SQRTPS128 struct {
}

func (v *SQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *SQRTPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SQRTPS128) Name() string {
	return "SQRTPS XMM (128 bit)"
}

func (v *SQRTPS128) Description() string {
	return "Compute square root of packed single-precision floats in XMM, lane-wise."
}

func (v *SQRTPS128) Stub() string {
	return stubSqrtps128
}

func (v *SQRTPS128) Assembly() string {
	return assemblySqrtps128
}

func (v *SQRTPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	sqrtps128(&floats, &ret)

	log.Printf("SQRTPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
