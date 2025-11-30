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
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVGATHERQPD256() *VGATHERQPD256 {
	return &VGATHERQPD256{
		base:  number.NewNamedFloatParameter("base", 512, 64),
		index: number.NewNamedIntParameter("index", 256, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 64, 16),
		src:   number.NewNamedFloatParameter("src", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VGATHERQPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VGATHERQPD256) Output() *number.Parameter {
	return v.ret
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

func (v *VGATHERQPD256) Run() {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	index := [4]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.src.FlatData()))

	vgatherqpd256(&base, &index, &mask, &ret)

	log.Printf("VGATHERQPD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VGATHERQPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
