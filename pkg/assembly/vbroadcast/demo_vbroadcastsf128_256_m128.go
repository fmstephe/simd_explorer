package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastsf128_256_m128.s
var assemblyVbroadcastsf128256M128 string

//go:embed stub_vbroadcastsf128_256_m128.go
var stubVbroadcastsf128256M128 string

type VBROADCASTSF128256M128 struct {
}

func (v *VBROADCASTSF128256M128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VBROADCASTSF128256M128) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VBROADCASTSF128256M128) Name() string {
	return "VBROADCASTSF128 (256 bit) m128"
}

func (v *VBROADCASTSF128256M128) Description() string {
	return "Broadcast 128-bit block (4x float32) from memory to both 128-bit lanes of YMM."
}

func (v *VBROADCASTSF128256M128) Stub() string {
	return stubVbroadcastsf128256M128
}

func (v *VBROADCASTSF128256M128) Assembly() string {
	return assemblyVbroadcastsf128256M128
}

func (v *VBROADCASTSF128256M128) Run(inputs [][]byte) (output []byte) {
	block := [4]float32{}
	copy(block[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vbroadcastsf128256M128(&block, &ret)

	log.Printf("VBROADCASTSF128256M128 block %v output %v", block, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VBROADCASTSF128256M128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
