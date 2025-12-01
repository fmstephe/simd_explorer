package vpgatherqq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherqq_256.s
var assemblyVpgatherqq256 string

//go:embed stub_vpgatherqq_256.go
var stubVpgatherqq256 string

type VPGATHERQQ256 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	src   *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERQQ256() *VPGATHERQQ256 {
	return &VPGATHERQQ256{
		base:  number.NewNamedUintParameter("base", 512, 64, 16),
		index: number.NewNamedIntParameter("index", 256, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 64, 16),
		src:   number.NewNamedUintParameter("src", 256, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPGATHERQQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.index,
		v.mask,
		v.src,
	}
}

func (v *VPGATHERQQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERQQ256) Name() string {
	return "VPGATHERQQ (256 bit) "
}

func (v *VPGATHERQQ256) Description() string {
	return "Gather 4 u64 elements from base + (i64 index * 8); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VPGATHERQQ256) Stub() string {
	return stubVpgatherqq256
}

func (v *VPGATHERQQ256) Assembly() string {
	return assemblyVpgatherqq256
}

func (v *VPGATHERQQ256) Run() {
	base := [8]uint64{}
	copy(base[:], number.ToUint64Slice(v.base.FlatData()))
	index := [4]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [4]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(v.src.FlatData()))

	vpgatherqq256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQQ256 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPGATHERQQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
