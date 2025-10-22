package subss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_subss_128.s
var assemblySubss128 string

//go:embed stub_subss_128.go
var stubSubss128 string

type SUBSS128 struct {
}

func (v *SUBSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SUBSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SUBSS128) Name() string {
	return "SUBSS (128 bit) "
}

func (v *SUBSS128) Description() string {
	return "TODO"
}

func (v *SUBSS128) Stub() string {
	return stubSubss128
}

func (v *SUBSS128) Assembly() string {
	return assemblySubss128
}

func (v *SUBSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	subss128(&floats1, &floats2, &ret)

	log.Printf("SUBSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SUBSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}