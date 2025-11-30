package vpgatherdd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherdd_128.s
var assemblyVpgatherdd128 string

//go:embed stub_vpgatherdd_128.go
var stubVpgatherdd128 string

type VPGATHERDD128 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERDD128() *VPGATHERDD128 {
	return &VPGATHERDD128{
		base:  number.NewNamedUintParameter("base", 256, 32, 16), // base memory (8 x u32)
		index: number.NewNamedIntParameter("index", 128, 32, 10), // indices (i32; lower 4 used)
		mask:  number.NewNamedUintParameter("mask", 128, 32, 16), // mask (MSB of each dword lane)
		src:   number.NewNamedUintParameter("src", 128, 32, 16),  // src (merge for masked-off lanes)
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPGATHERDD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VPGATHERDD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERDD128) Name() string {
	return "VPGATHERDD (128 bit) "
}

func (v *VPGATHERDD128) Description() string {
	return "Gather 4 u32 elements from base + (i32 index * 4); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VPGATHERDD128) Stub() string {
	return stubVpgatherdd128
}

func (v *VPGATHERDD128) Assembly() string {
	return assemblyVpgatherdd128
}

func (v *VPGATHERDD128) Run() {
	base := [8]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(v.src.FlatData()))

	vpgatherdd128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	retSlc := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VPGATHERDD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
