package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastss_128_m32.s
var assemblyVbroadcastss128M32 string

//go:embed stub_vbroadcastss_128_m32.go
var stubVbroadcastss128M32 string

type VBROADCASTSS128M32 struct {
}

func (v *VBROADCASTSS128M32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(32, 32),
	}
}

func (v *VBROADCASTSS128M32) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VBROADCASTSS128M32) Name() string {
	return "VBROADCASTSS (128 bit) m32"
}

func (v *VBROADCASTSS128M32) Description() string {
	return "Broadcast 32-bit scalar from memory into all 4 lanes of XMM."
}

func (v *VBROADCASTSS128M32) Stub() string {
	return stubVbroadcastss128M32
}

func (v *VBROADCASTSS128M32) Assembly() string {
	return assemblyVbroadcastss128M32
}

func (v *VBROADCASTSS128M32) Run(inputs [][]byte) (output []byte) {
	scalar := number.ToFloat32(inputs[0])

	ret := [4]float32{}

	vbroadcastss128M32(&scalar, &ret)

	log.Printf("VBROADCASTSS128M32 scalar %v output %v", scalar, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VBROADCASTSS128M32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
