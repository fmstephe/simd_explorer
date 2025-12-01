package vpgatherdq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherdq_128.s
var assemblyVpgatherdq128 string

//go:embed stub_vpgatherdq_128.go
var stubVpgatherdq128 string

type VPGATHERDQ128 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERDQ128() *VPGATHERDQ128 {
	return &VPGATHERDQ128{
		base:  number.NewNamedUintParameter("base", 256, 64, 10),
		index: number.NewNamedIntParameter("index", 128, 32, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPGATHERDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
	}
}

func (v *VPGATHERDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERDQ128) Name() string {
	return "VPGATHERDQ (128 bit) "
}

func (v *VPGATHERDQ128) Description() string {
	return "Gather 2 u32 elements from base + (i64 index * 4); lanes selected by mask (MSB). Lower 2 indices are used; result merged into src."
}

func (v *VPGATHERDQ128) Stub() string {
	return stubVpgatherdq128
}

func (v *VPGATHERDQ128) Assembly() string {
	return assemblyVpgatherdq128
}

func (v *VPGATHERDQ128) Run() {
	base := [4]uint64{}
	copy(base[:], number.ToUint64Slice(v.base.FlatData()))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	mask := [2]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpgatherdq128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDQ128 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPGATHERDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
