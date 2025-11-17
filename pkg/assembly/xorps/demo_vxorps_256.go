package xorps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vxorps_256.s
var assemblyVxorps256 string

//go:embed stub_vxorps_256.go
var stubVxorps256 string

type VXORPS256 struct {
}

func (v *VXORPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VXORPS256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VXORPS256) Name() string {
	return "VXORPS (256 bit)"
}

func (v *VXORPS256) Description() string {
	return "Bitwise XOR of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VXORPS256) Stub() string {
	return stubVxorps256
}

func (v *VXORPS256) Assembly() string {
	return assemblyVxorps256
}

func (v *VXORPS256) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vxorps256(&vals1, &vals2, &ret)

	log.Printf("VXORPS256 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VXORPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
