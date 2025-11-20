package vbroadcast

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcastsd_256_m64.s
var assemblyVbroadcastsd256M64 string

//go:embed stub_vbroadcastsd_256_m64.go
var stubVbroadcastsd256M64 string

type VBROADCASTSD256M64 struct {
}

func (v *VBROADCASTSD256M64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 64),
	}
}

func (v *VBROADCASTSD256M64) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64)
}

func (v *VBROADCASTSD256M64) Name() string {
	return "VBROADCASTSD (256 bit) m64"
}

func (v *VBROADCASTSD256M64) Description() string {
	return "Broadcast 64-bit scalar from memory into all 4 lanes of YMM."
}

func (v *VBROADCASTSD256M64) Stub() string {
	return stubVbroadcastsd256M64
}

func (v *VBROADCASTSD256M64) Assembly() string {
	return assemblyVbroadcastsd256M64
}

func (v *VBROADCASTSD256M64) Run(inputs [][]byte) (output []byte) {
	scalar := number.ToFloat64(inputs[0])

	ret := [4]float64{}

	vbroadcastsd256M64(&scalar, &ret)

	log.Printf("VBROADCASTSD256M64 scalar %v output %v", scalar, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VBROADCASTSD256M64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
