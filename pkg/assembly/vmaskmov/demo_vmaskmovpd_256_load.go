package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_256_load.s
var assemblyVmaskmovpd256Load string

//go:embed stub_vmaskmovpd_256_load.go
var stubVmaskmovpd256Load string

type VMASKMOVPD256LOAD struct {
}

func (v *VMASKMOVPD256LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),    // memory
		number.NewUintParameter(256, 64, 16), // mask
	}
}

func (v *VMASKMOVPD256LOAD) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64) // loaded vector
}

func (v *VMASKMOVPD256LOAD) Name() string {
	return "VMASKMOVPD (256 bit) load"
}

func (v *VMASKMOVPD256LOAD) Description() string {
	return "Masked load of packed double-precision floats (per 128-bit lane); unselected lanes are zeroed."
}

func (v *VMASKMOVPD256LOAD) Stub() string {
	return stubVmaskmovpd256Load
}

func (v *VMASKMOVPD256LOAD) Assembly() string {
	return assemblyVmaskmovpd256Load
}

func (v *VMASKMOVPD256LOAD) Run(inputs [][]byte) (output []byte) {
	mem := [4]float64{}
	copy(mem[:], number.ToFloat64Slice(inputs[0]))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[1]))

	ret := [4]float64{}

	vmaskmovpd256Load(&mem, &mask, &ret)

	log.Printf("VMASKMOVPD256LOAD mem %v mask %v ret %v", mem, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VMASKMOVPD256LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
