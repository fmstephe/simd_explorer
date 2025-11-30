package vpgatherdq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherdq_256.s
var assemblyVpgatherdq256 string

//go:embed stub_vpgatherdq_256.go
var stubVpgatherdq256 string

type VPGATHERDQ256 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERDQ256() *VPGATHERDQ256 {
	return &VPGATHERDQ256{
		base:  number.NewNamedUintParameter("base", 512, 64, 10),
		index: number.NewNamedIntParameter("index", 128, 32, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPGATHERDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{v.base, v.index, v.mask}
}

func (v *VPGATHERDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERDQ256) Name() string {
	return "VPGATHERDQ (256 bit) "
}

func (v *VPGATHERDQ256) Description() string {
	return "Gather 4 u32 elements from base + (i64 index * 4); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VPGATHERDQ256) Stub() string {
	return stubVpgatherdq256
}

func (v *VPGATHERDQ256) Assembly() string {
	return assemblyVpgatherdq256
}

func (v *VPGATHERDQ256) Run() {
	base := [8]uint64{}
	copy(base[:], number.ToUint64Slice(v.base.FlatData()))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(v.index.FlatData()))
	mask := [4]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpgatherdq256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDQ256 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPGATHERDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
