package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpps_128.s
var assemblyVrcpps128 string

//go:embed stub_vrcpps_128.go
var stubVrcpps128 string

type VRCPPS128 struct {
}

func (v *VRCPPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VRCPPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VRCPPS128) Name() string {
	return "VRCPPS XMM (128 bit)"
}

func (v *VRCPPS128) Description() string {
	return "TODO"
}

func (v *VRCPPS128) Stub() string {
	return stubVrcpps128
}

func (v *VRCPPS128) Assembly() string {
	return assemblyVrcpps128
}

func (v *VRCPPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vrcpps128(&floats, &ret)

	log.Printf("VRCPPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VRCPPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
