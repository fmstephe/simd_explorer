package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsubps_256.s
var assemblyVsubps256 string

//go:embed stub_vsubps_256.go
var stubVsubps256 string

type VSUBPS256 struct {
}

func (v *VSUBPS256) Inputs() []*number.Parameter {
    return []*number.Parameter{
        number.NewFloatParameter(256, 32),
        number.NewFloatParameter(256, 32),
    }
}

func (v *VSUBPS256) Output() *number.Parameter {
    return number.NewFloatParameter(256, 32)
}

func (v *VSUBPS256) Name() string {
	return "VSUBPS (256 bit) "
}

func (v *VSUBPS256) Description() string {
	return "TODO"
}

func (v *VSUBPS256) Stub() string {
	return stubVsubps256
}

func (v *VSUBPS256) Assembly() string {
	return assemblyVsubps256
}

func (v *VSUBPS256) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
    floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

    vsubps256(&floats1, &floats2, &ret)

	log.Printf("VSUBPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSUBPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
