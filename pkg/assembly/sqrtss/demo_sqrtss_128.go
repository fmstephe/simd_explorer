package sqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_sqrtss_128.s
var assemblySqrtss128 string

//go:embed stub_sqrtss_128.go
var stubSqrtss128 string

type SQRTSS128 struct {
}

func (v *SQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *SQRTSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SQRTSS128) Name() string {
	return "SQRTSS (128 bit) "
}

func (v *SQRTSS128) Description() string {
	return "TODO"
}

func (v *SQRTSS128) Stub() string {
	return stubSqrtss128
}

func (v *SQRTSS128) Assembly() string {
	return assemblySqrtss128
}

func (v *SQRTSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	sqrtss128(&floats, &ret)

	log.Printf("SQRTSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}