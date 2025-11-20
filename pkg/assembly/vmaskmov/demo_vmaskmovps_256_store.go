package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_256_store.s
var assemblyVmaskmovps256Store string

//go:embed stub_vmaskmovps_256_store.go
var stubVmaskmovps256Store string

type VMASKMOVPS256STORE struct {
}

func (v *VMASKMOVPS256STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),    // src
		number.NewUintParameter(256, 32, 16), // mask
	}
}

func (v *VMASKMOVPS256STORE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32) // memory after store
}

func (v *VMASKMOVPS256STORE) Name() string {
	return "VMASKMOVPS (256 bit) store"
}

func (v *VMASKMOVPS256STORE) Description() string {
	return "Masked store of packed single-precision floats (per 128-bit lane)."
}

func (v *VMASKMOVPS256STORE) Stub() string {
	return stubVmaskmovps256Store
}

func (v *VMASKMOVPS256STORE) Assembly() string {
	return assemblyVmaskmovps256Store
}

func (v *VMASKMOVPS256STORE) Run(inputs [][]byte) (output []byte) {
	src := [8]float32{}
	copy(src[:], number.ToFloat32Slice(inputs[0]))
	mask := [8]float32{}
	copy(mask[:], number.ToFloat32Slice(inputs[1]))

	mem := [8]float32{}

	vmaskmovps256Store(&src, &mask, &mem)

	log.Printf("VMASKMOVPS256STORE src %v mask %v mem %v", src, mask, mem)

	return number.Float32SliceToBytes(mem[:])
}

func (v *VMASKMOVPS256STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
