package xorps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_xorps_128.s
var assemblyXorps128 string

//go:embed stub_xorps_128.go
var stubXorps128 string

type XORPS128 struct {
}

func (v *XORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *XORPS128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
}

func (v *XORPS128) Name() string {
	return "XORPS (128 bit)"
}

func (v *XORPS128) Description() string {
	return "Bitwise XOR of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *XORPS128) Stub() string {
	return stubXorps128
}

func (v *XORPS128) Assembly() string {
	return assemblyXorps128
}

func (v *XORPS128) Run(inputs [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(inputs[0]))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	xorps128(&vals1, &vals2, &ret)

	log.Printf("XORPS128 input %v %v output %v", vals1, vals2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *XORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
