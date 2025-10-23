package divss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_divss_128.s
var assemblyDivss128 string

//go:embed stub_divss_128.go
var stubDivss128 string

type DIVSS128 struct {
}

func (v *DIVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *DIVSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *DIVSS128) Name() string {
	return "DIVSS (128 bit) "
}

func (v *DIVSS128) Description() string {
	return "TODO"
}

func (v *DIVSS128) Stub() string {
	return stubDivss128
}

func (v *DIVSS128) Assembly() string {
	return assemblyDivss128
}

func (v *DIVSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	divss128(&floats1, &floats2, &ret)

	log.Printf("DIVSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *DIVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
