package vgatherdpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vgatherdpd_128.s
var assemblyVgatherdpd128 string

//go:embed stub_vgatherdpd_128.go
var stubVgatherdpd128 string

type VGATHERDPD128 struct {
}

func (v *VGATHERDPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(512, 64),    // base memory (8 x f64)
		number.NewUintParameter(128, 32, 10), // indices (i32; lower 2 used)
		number.NewUintParameter(128, 64, 16), // mask (MSB of each f64 lane)
		number.NewFloatParameter(128, 64),    // src/dst (merge for masked-off lanes)
	}
}

func (v *VGATHERDPD128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64) // gathered vector
}

func (v *VGATHERDPD128) Name() string {
	return "VGATHERDPD (128 bit) "
}

func (v *VGATHERDPD128) Description() string {
	return "Gather 2 f64 elements from base + (i32 index * 8); lanes selected by mask (MSB). Lower 2 indices are used; result merged into src."
}

func (v *VGATHERDPD128) Stub() string {
	return stubVgatherdpd128
}

func (v *VGATHERDPD128) Assembly() string {
	return assemblyVgatherdpd128
}

func (v *VGATHERDPD128) Run(inputs [][]byte) (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(inputs[0]))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1]))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[2]))

	ret := [2]float64{}
	// seed destination with src for masked merge behaviour
	copy(ret[:], number.ToFloat64Slice(inputs[3]))

	vgatherdpd128(&base, &index, &mask, &ret)

	log.Printf("VGATHERDPD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VGATHERDPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
