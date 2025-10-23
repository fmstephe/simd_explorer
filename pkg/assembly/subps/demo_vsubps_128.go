package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsubps_128.s
var assemblyVsubps128 string

//go:embed stub_vsubps_128.go
var stubVsubps128 string

type VSUBPS128 struct {
}

func (v *VSUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSUBPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSUBPS128) Name() string {
	return "VSUBPS (128 bit) "
}

func (v *VSUBPS128) Description() string {
	return "TODO"
}

func (v *VSUBPS128) Stub() string {
	return stubVsubps128
}

func (v *VSUBPS128) Assembly() string {
	return assemblyVsubps128
}

func (v *VSUBPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vsubps128(&floats1, &floats2, &ret)

	log.Printf("VSUBPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
