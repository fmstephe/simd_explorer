package subss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsubss_128.s
var assemblyVsubss128 string

//go:embed stub_vsubss_128.go
var stubVsubss128 string

type VSUBSS128 struct {
}

func (v *VSUBSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSUBSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSUBSS128) Name() string {
	return "VSUBSS (128 bit) "
}

func (v *VSUBSS128) Description() string {
	return "TODO"
}

func (v *VSUBSS128) Stub() string {
	return stubVsubss128
}

func (v *VSUBSS128) Assembly() string {
	return assemblyVsubss128
}

func (v *VSUBSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vsubss128(&floats1, &floats2, &ret)

	log.Printf("VSUBSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSUBSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}