package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_128_load.s
var assemblyVmaskmovps128Load string

//go:embed stub_vmaskmovps_128_load.go
var stubVmaskmovps128Load string

type VMASKMOVPS128LOAD struct {
}

func (v *VMASKMOVPS128LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),    // memory
		number.NewUintParameter(128, 32, 16), // mask
	}
}

func (v *VMASKMOVPS128LOAD) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32) // loaded vector
}

func (v *VMASKMOVPS128LOAD) Name() string {
	return "VMASKMOVPS (128 bit) load"
}

func (v *VMASKMOVPS128LOAD) Description() string {
	return "Masked load of packed single-precision floats; unselected lanes are zeroed."
}

func (v *VMASKMOVPS128LOAD) Stub() string {
	return stubVmaskmovps128Load
}

func (v *VMASKMOVPS128LOAD) Assembly() string {
	return assemblyVmaskmovps128Load
}

func (v *VMASKMOVPS128LOAD) Run(inputs [][]byte) (output []byte) {
	mem := [4]float32{}
	copy(mem[:], number.ToFloat32Slice(inputs[0]))
	mask := [4]float32{}
	copy(mask[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmaskmovps128Load(&mem, &mask, &ret)

	log.Printf("VMASKMOVPS128LOAD mem %v mask %v ret %v", mem, mask, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMASKMOVPS128LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
