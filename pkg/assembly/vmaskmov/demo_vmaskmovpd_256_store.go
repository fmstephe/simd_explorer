package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_256_store.s
var assemblyVmaskmovpd256Store string

//go:embed stub_vmaskmovpd_256_store.go
var stubVmaskmovpd256Store string

type VMASKMOVPD256STORE struct {
}

func (v *VMASKMOVPD256STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),    // src (4 lanes)
		number.NewUintParameter(256, 64, 16), // mask
	}
}

func (v *VMASKMOVPD256STORE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64) // memory after store
}

func (v *VMASKMOVPD256STORE) Name() string {
	return "VMASKMOVPD (256 bit) store"
}

func (v *VMASKMOVPD256STORE) Description() string {
	return "Masked store of packed double-precision floats (per 128-bit lane)."
}

func (v *VMASKMOVPD256STORE) Stub() string {
	return stubVmaskmovpd256Store
}

func (v *VMASKMOVPD256STORE) Assembly() string {
	return assemblyVmaskmovpd256Store
}

func (v *VMASKMOVPD256STORE) Run(inputs [][]byte) (output []byte) {
	src := [4]float64{}
	copy(src[:], number.ToFloat64Slice(inputs[0]))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[1]))

	mem := [4]float64{}

	vmaskmovpd256Store(&src, &mask, &mem)

	log.Printf("VMASKMOVPD256STORE src %v mask %v mem %v", src, mask, mem)

	return number.Float64SliceToBytes(mem[:])
}

func (v *VMASKMOVPD256STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
