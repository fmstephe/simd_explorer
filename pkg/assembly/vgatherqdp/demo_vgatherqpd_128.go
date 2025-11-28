package vgatherqdp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vgatherqpd_128.s
var assemblyVgatherqpd128 string

//go:embed stub_vgatherqpd_128.go
var stubVgatherqpd128 string

type VGATHERQPD128 struct {
}

func (v *VGATHERQPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(512, 64),    // base memory (8 x f64)
		number.NewIntParameter(128, 64, 10),  // indices (i64; lower 2 used)
		number.NewUintParameter(128, 64, 16), // mask (MSB of each f64 lane)
		number.NewFloatParameter(128, 64),    // src/dst (merge for masked-off lanes)
	}
}

func (v *VGATHERQPD128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64) // gathered vector
}

func (v *VGATHERQPD128) Name() string {
	return "VGATHERQPD (128 bit) "
}

func (v *VGATHERQPD128) Description() string {
	return "Gather 2 f64 elements from base + (i64 index * 8); lanes selected by mask (MSB). Lower 2 indices are used; result merged into src."
}

func (v *VGATHERQPD128) Stub() string {
	return stubVgatherqpd128
}

func (v *VGATHERQPD128) Assembly() string {
	return assemblyVgatherqpd128
}

func (v *VGATHERQPD128) Run(inputs [][]byte) (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(inputs[0]))
	index := [2]uint64{}
	copy(index[:], number.ToUint64Slice(inputs[1]))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(inputs[2]))

	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(inputs[3]))

	vgatherqpd128(&base, &index, &mask, &ret)

	log.Printf("VGATHERQPD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VGATHERQPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
