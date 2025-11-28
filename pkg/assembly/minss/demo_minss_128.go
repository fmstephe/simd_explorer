package minss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_minss_128.s
var assemblyMinss128 string

//go:embed stub_minss_128.go
var stubMinss128 string

type MINSS128 struct {
}

func (v *MINSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MINSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MINSS128) Name() string {
	return "MINSS (128 bit) "
}

func (v *MINSS128) Description() string {
	return "Compute minimum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *MINSS128) Stub() string {
	return stubMinss128
}

func (v *MINSS128) Assembly() string {
	return assemblyMinss128
}

func (v *MINSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	minss128(&floats1, &floats2, &ret)

	log.Printf("MINSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MINSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
