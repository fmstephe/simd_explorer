package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_128_store.s
var assemblyVmaskmovpd128Store string

//go:embed stub_vmaskmovpd_128_store.go
var stubVmaskmovpd128Store string

type VMASKMOVPD128STORE struct {
}

func (v *VMASKMOVPD128STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),    // src (2 lanes)
		number.NewUintParameter(128, 64, 16), // mask
	}
}

func (v *VMASKMOVPD128STORE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64) // memory after store
}

func (v *VMASKMOVPD128STORE) Name() string {
	return "VMASKMOVPD (128 bit) store"
}

func (v *VMASKMOVPD128STORE) Description() string {
	return "Masked store of packed double-precision floats: store lanes where mask sign-bit is set."
}

func (v *VMASKMOVPD128STORE) Stub() string {
	return stubVmaskmovpd128Store
}

func (v *VMASKMOVPD128STORE) Assembly() string {
	return assemblyVmaskmovpd128Store
}

func (v *VMASKMOVPD128STORE) Run(inputs [][]byte) (output []byte) {
	src := [2]float64{}
	copy(src[:], number.ToFloat64Slice(inputs[0]))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[1]))

	mem := [2]float64{}

	vmaskmovpd128Store(&src, &mask, &mem)

	log.Printf("VMASKMOVPD128STORE src %v mask %v mem %v", src, mask, mem)

	return number.Float64SliceToBytes(mem[:])
}

func (v *VMASKMOVPD128STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
