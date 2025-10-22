package rcpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpss_128.s
var assemblyVrcpss128 string

//go:embed stub_vrcpss_128.go
var stubVrcpss128 string

type VRCPSS128 struct {
}

func (v *VRCPSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VRCPSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VRCPSS128) Name() string {
	return "VRCPSS (128 bit) "
}

func (v *VRCPSS128) Description() string {
	return "TODO"
}

func (v *VRCPSS128) Stub() string {
	return stubVrcpss128
}

func (v *VRCPSS128) Assembly() string {
	return assemblyVrcpss128
}

func (v *VRCPSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vrcpss128(&floats, &ret)

	log.Printf("VRCPSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VRCPSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}