package maxss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_maxss_128.s
var assemblyMaxss128 string

//go:embed stub_maxss_128.go
var stubMaxss128 string

type MAXSS128 struct {
}

func (v *MAXSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MAXSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MAXSS128) Name() string {
	return "MAXSS (128 bit) "
}

func (v *MAXSS128) Description() string {
	return "Compute maximum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *MAXSS128) Stub() string {
	return stubMaxss128
}

func (v *MAXSS128) Assembly() string {
	return assemblyMaxss128
}

func (v *MAXSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	maxss128(&floats1, &floats2, &ret)

	log.Printf("MAXSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MAXSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
