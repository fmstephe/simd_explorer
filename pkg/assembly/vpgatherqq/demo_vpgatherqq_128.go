package vpgatherqq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherqq_128.s
var assemblyVpgatherqq128 string

//go:embed stub_vpgatherqq_128.go
var stubVpgatherqq128 string

type VPGATHERQQ128 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERQQ128() *VPGATHERQQ128 {
	return &VPGATHERQQ128{
		base:  number.NewNamedUintParameter("base", 512, 64, 16),
		index: number.NewNamedIntParameter("index", 128, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 64, 16),
		src:   number.NewNamedUintParameter("src", 128, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 16),
	}
}

func (v *VPGATHERQQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.base, v.index, v.mask, v.src}
}

func (v *VPGATHERQQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERQQ128) Name() string {
	return "VPGATHERQQ (128 bit) "
}

func (v *VPGATHERQQ128) Description() string {
	return "Gather 2 u64 elements from base + (i64 index * 8); lanes selected by mask (MSB). Lower 2 indices are used; result merged into src."
}

func (v *VPGATHERQQ128) Stub() string {
	return stubVpgatherqq128
}

func (v *VPGATHERQQ128) Assembly() string {
	return assemblyVpgatherqq128
}

func (v *VPGATHERQQ128) Run() {
	base := [8]uint64{}
	copy(base[:], number.ToUint64Slice(v.base.FlatData()))
	index := [2]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [2]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(v.src.FlatData()))

	vpgatherqq128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQQ128 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPGATHERQQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
