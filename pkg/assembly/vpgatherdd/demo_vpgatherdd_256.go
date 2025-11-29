package vpgatherdd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherdd_256.s
var assemblyVpgatherdd256 string

//go:embed stub_vpgatherdd_256.go
var stubVpgatherdd256 string

type VPGATHERDD256 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERDD256() *VPGATHERDD256 {
	return &VPGATHERDD256{
		base:  number.NewNamedUintParameter("base", 512, 32, 16), // base memory (16 x u32)
		index: number.NewNamedIntParameter("index", 256, 32, 10), // indices (i32; lower 8 used)
		mask:  number.NewNamedUintParameter("mask", 256, 32, 16), // mask (MSB of each dword lane)
		src:   number.NewNamedUintParameter("src", 256, 32, 16),  // src (merge for masked-off lanes)
		ret:   number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VPGATHERDD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VPGATHERDD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERDD256) Name() string {
	return "VPGATHERDD (256 bit) "
}

func (v *VPGATHERDD256) Description() string {
	return "Gather 8 u32 elements from base + (i32 index * 4); lanes selected by mask (MSB). Lower 8 indices are used; result merged into src."
}

func (v *VPGATHERDD256) Stub() string {
	return stubVpgatherdd256
}

func (v *VPGATHERDD256) Assembly() string {
	return assemblyVpgatherdd256
}

func (v *VPGATHERDD256) Run(inputs [][]byte) (output []byte) {
	base := [16]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	index := [8]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	mask := [8]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))
	ret := [8]uint32{}
	copy(ret[:], number.ToUint32Slice(v.src.FlatData()))

	vpgatherdd256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	retSlc := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VPGATHERDD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
