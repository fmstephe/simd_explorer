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
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVGATHERDPD256() *VGATHERDPD256 {
	return &VGATHERDPD256{
		base:  number.NewNamedFloatParameter("base", 512, 64),
		index: number.NewNamedIntParameter("index", 128, 32, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 64, 16),
		src:   number.NewNamedFloatParameter("src", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VGATHERDPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VGATHERDPD256) Output() *number.Parameter {
	return v.ret
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

func (v *VGATHERDPD256) Run() {
	base := [8]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	// indices are provided as signed i32 in base 10; we read bits from Parameter storage
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	// mask provided in hex; bits are interpreted as f64 lane masks (MSB)
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [4]float64{}
	// seed destination with src for masked merge behaviour
	copy(ret[:], number.ToFloat64Slice(v.src.FlatData()))

	vgatherdpd256(&base, &index, &mask, &ret)

	log.Printf("VGATHERDPD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VGATHERDPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
