package addss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_addss_128.s
var assemblyAddss128 string

//go:embed stub_addss_128.go
var stubAddss128 string

type ADDSS128 struct {
}

func (v *ADDSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *ADDSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *ADDSS128) Name() string {
	return "ADDSS (128 bit) "
}

func (v *ADDSS128) Description() string {
	return "TODO"
}

func (v *ADDSS128) Stub() string {
	return stubAddss128
}

func (v *ADDSS128) Assembly() string {
	return assemblyAddss128
}

func (v *ADDSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	addss128(&floats1, &floats2, &ret)

	log.Printf("ADDSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *ADDSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
