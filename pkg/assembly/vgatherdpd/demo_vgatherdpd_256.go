package vgatherdpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vgatherdpd_256.s
var assemblyVgatherdpd256 string

//go:embed stub_vgatherdpd_256.go
var stubVgatherdpd256 string

type VGATHERDPD256 struct {
}

func (v *VGATHERDPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(512, 64),    // base memory (8 x f64)
		number.NewIntParameter(128, 32, 10),  // indices (i32; lower 4 used)
		number.NewUintParameter(256, 64, 16), // mask (MSB of each f64 lane)
		number.NewFloatParameter(256, 64),    // src/dst (merge for masked-off lanes)
	}
}

func (v *VGATHERDPD256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64) // gathered vector
}

func (v *VGATHERDPD256) Name() string {
	return "VGATHERDPD (256 bit) "
}

func (v *VGATHERDPD256) Description() string {
	return "Gather 4 f64 elements from base + (i32 index * 8); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VGATHERDPD256) Stub() string {
	return stubVgatherdpd256
}

func (v *VGATHERDPD256) Assembly() string {
	return assemblyVgatherdpd256
}

func (v *VGATHERDPD256) Run(inputs [][]byte) (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(inputs[0]))
	// indices are provided as signed i32 in base 10 in the UI; we read bits
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1]))
	// mask provided in hex; bits are interpreted as f64 lane masks (MSB)
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[2]))

	ret := [4]float64{}
	// seed destination with src for masked merge behaviour
	copy(ret[:], number.ToFloat64Slice(inputs[3]))

	vgatherdpd256(&base, &index, &mask, &ret)

	log.Printf("VGATHERDPD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VGATHERDPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
