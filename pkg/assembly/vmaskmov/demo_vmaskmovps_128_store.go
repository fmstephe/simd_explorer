package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_128_store.s
var assemblyVmaskmovps128Store string

//go:embed stub_vmaskmovps_128_store.go
var stubVmaskmovps128Store string

type VMASKMOVPS128STORE struct {
}

func (v *VMASKMOVPS128STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),    // src
		number.NewUintParameter(128, 32, 16), // mask (sign bit selects, 16)
	}
}

func (v *VMASKMOVPS128STORE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32) // memory after store
}

func (v *VMASKMOVPS128STORE) Name() string {
	return "VMASKMOVPS (128 bit) store"
}

func (v *VMASKMOVPS128STORE) Description() string {
	return "Masked store of packed single-precision floats: store lanes where mask sign-bit is set."
}

func (v *VMASKMOVPS128STORE) Stub() string {
	return stubVmaskmovps128Store
}

func (v *VMASKMOVPS128STORE) Assembly() string {
	return assemblyVmaskmovps128Store
}

func (v *VMASKMOVPS128STORE) Run(inputs [][]byte) (output []byte) {
	src := [4]float32{}
	copy(src[:], number.ToFloat32Slice(inputs[0]))
	mask := [4]float32{}
	copy(mask[:], number.ToFloat32Slice(inputs[1]))

	mem := [4]float32{}

	vmaskmovps128Store(&src, &mask, &mem)

	log.Printf("VMASKMOVPS128STORE src %v mask %v mem %v", src, mask, mem)

	return number.Float32SliceToBytes(mem[:])
}

func (v *VMASKMOVPS128STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
