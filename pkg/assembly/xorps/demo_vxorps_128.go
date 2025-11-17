package xorps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vxorps_128.s
var assemblyVxorps128 string

//go:embed stub_vxorps_128.go
var stubVxorps128 string

type VXORPS128 struct {
}

func (v *VXORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VXORPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *VXORPS128) Name() string {
	return "VXORPS (128 bit)"
}

func (v *VXORPS128) Description() string {
	return "Bitwise XOR of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VXORPS128) Stub() string {
	return stubVxorps128
}

func (v *VXORPS128) Assembly() string {
	return assemblyVxorps128
}

func (v *VXORPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vxorps128(&vals1, &vals2, &ret)

	log.Printf("VXORPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VXORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
