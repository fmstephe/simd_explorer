package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_128_load.s
var assemblyVmaskmovpd128Load string

//go:embed stub_vmaskmovpd_128_load.go
var stubVmaskmovpd128Load string

type VMASKMOVPD128LOAD struct {
}

func (v *VMASKMOVPD128LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),    // memory
		number.NewUintParameter(128, 64, 16), // mask
	}
}

func (v *VMASKMOVPD128LOAD) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64) // loaded vector
}

func (v *VMASKMOVPD128LOAD) Name() string {
	return "VMASKMOVPD (128 bit) load"
}

func (v *VMASKMOVPD128LOAD) Description() string {
	return "Masked load of packed double-precision floats; unselected lanes are zeroed."
}

func (v *VMASKMOVPD128LOAD) Stub() string {
	return stubVmaskmovpd128Load
}

func (v *VMASKMOVPD128LOAD) Assembly() string {
	return assemblyVmaskmovpd128Load
}

func (v *VMASKMOVPD128LOAD) Run(inputs [][]byte) (output []byte) {
	mem := [2]float64{}
	copy(mem[:], number.ToFloat64Slice(inputs[0]))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[1]))

	ret := [2]float64{}

	vmaskmovpd128Load(&mem, &mask, &ret)

	log.Printf("VMASKMOVPD128LOAD mem %v mask %v ret %v", mem, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VMASKMOVPD128LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
