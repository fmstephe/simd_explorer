package vgatherqdp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vgatherqpd_256.s
var assemblyVgatherqpd256 string

//go:embed stub_vgatherqpd_256.go
var stubVgatherqpd256 string

type VGATHERQPD256 struct {
}

func (v *VGATHERQPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(512, 64),    // base memory (8 x f64)
		number.NewIntParameter(256, 64, 10),  // indices (i64; lower 4 used)
		number.NewUintParameter(256, 64, 16), // mask (MSB of each f64 lane)
		number.NewFloatParameter(256, 64),    // src/dst (merge for masked-off lanes)
	}
}

func (v *VGATHERQPD256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64) // gathered vector
}

func (v *VGATHERQPD256) Name() string {
	return "VGATHERQPD (256 bit) "
}

func (v *VGATHERQPD256) Description() string {
	return "Gather 4 f64 elements from base + (i64 index * 8); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VGATHERQPD256) Stub() string {
	return stubVgatherqpd256
}

func (v *VGATHERQPD256) Assembly() string {
	return assemblyVgatherqpd256
}

func (v *VGATHERQPD256) Run(inputs [][]byte) (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(inputs[0]))
	index := [4]uint64{}
	copy(index[:], number.ToUint64Slice(inputs[1]))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[2]))

	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(inputs[3]))

	vgatherqpd256(&base, &index, &mask, &ret)

	log.Printf("VGATHERQPD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VGATHERQPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
