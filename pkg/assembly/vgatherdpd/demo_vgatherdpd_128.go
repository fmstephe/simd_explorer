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
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVGATHERDPD128() *VGATHERDPD128 {
	return &VGATHERDPD128{
		base:  number.NewNamedFloatParameter("base", 512, 64),
		index: number.NewNamedIntParameter("index", 128, 32, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 64, 16),
		src:   number.NewNamedFloatParameter("src", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VGATHERDPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VGATHERDPD128) Output() *number.Parameter {
	return v.ret
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

func (v *VGATHERDPD128) Run() (output []byte) {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [2]float64{}
	// seed destination with src for masked merge behaviour
	copy(ret[:], number.ToFloat64Slice(v.src.FlatData()))

	vgatherdpd128(&base, &index, &mask, &ret)

	log.Printf("VGATHERDPD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VGATHERDPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
