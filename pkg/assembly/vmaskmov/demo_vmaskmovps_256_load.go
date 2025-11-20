package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_256_load.s
var assemblyVmaskmovps256Load string

//go:embed stub_vmaskmovps_256_load.go
var stubVmaskmovps256Load string

type VMASKMOVPS256LOAD struct {
}

func (v *VMASKMOVPS256LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),    // memory
		number.NewUintParameter(256, 32, 16), // mask
	}
}

func (v *VMASKMOVPS256LOAD) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32) // loaded vector
}

func (v *VMASKMOVPS256LOAD) Name() string {
	return "VMASKMOVPS (256 bit) load"
}

func (v *VMASKMOVPS256LOAD) Description() string {
	return "Masked load of packed single-precision floats (per 128-bit lane); unselected lanes are zeroed."
}

func (v *VMASKMOVPS256LOAD) Stub() string {
	return stubVmaskmovps256Load
}

func (v *VMASKMOVPS256LOAD) Assembly() string {
	return assemblyVmaskmovps256Load
}

func (v *VMASKMOVPS256LOAD) Run(inputs [][]byte) (output []byte) {
	mem := [8]float32{}
	copy(mem[:], number.ToFloat32Slice(inputs[0]))
	mask := [8]float32{}
	copy(mask[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vmaskmovps256Load(&mem, &mask, &ret)

	log.Printf("VMASKMOVPS256LOAD mem %v mask %v ret %v", mem, mask, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMASKMOVPS256LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
