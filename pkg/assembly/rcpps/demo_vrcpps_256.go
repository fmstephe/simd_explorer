package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpps_256.s
var assemblyVrcpps256 string

//go:embed stub_vrcpps_256.go
var stubVrcpps256 string

type VRCPPS256 struct {
}

func (v *VRCPPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VRCPPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VRCPPS256) Name() string {
	return "VRCPPS YMM (256 bit)"
}

func (v *VRCPPS256) Description() string {
	return "TODO"
}

func (v *VRCPPS256) Stub() string {
	return stubVrcpps256
}

func (v *VRCPPS256) Assembly() string {
	return assemblyVrcpps256
}

func (v *VRCPPS256) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vrcpps256(&floats, &ret)

	log.Printf("VRCPPS256 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VRCPPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
