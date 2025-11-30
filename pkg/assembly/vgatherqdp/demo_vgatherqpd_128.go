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
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVGATHERQPD128() *VGATHERQPD128 {
	return &VGATHERQPD128{
		base:  number.NewNamedFloatParameter("base", 512, 64),
		index: number.NewNamedIntParameter("index", 128, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 64, 16),
		src:   number.NewNamedFloatParameter("src", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VGATHERQPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VGATHERQPD128) Output() *number.Parameter {
	return v.ret
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

func (v *VGATHERQPD128) Run() (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	index := [2]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.src.FlatData()))

	vgatherqpd128(&base, &index, &mask, &ret)

	log.Printf("VGATHERQPD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VGATHERQPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
