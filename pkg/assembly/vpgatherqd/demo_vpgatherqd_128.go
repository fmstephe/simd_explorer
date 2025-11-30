package vpgatherqd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherqd_128.s
var assemblyVpgatherqd128 string

//go:embed stub_vpgatherqd_128.go
var stubVpgatherqd128 string

type VPGATHERQD128 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERQD128() *VPGATHERQD128 {
	return &VPGATHERQD128{
		base:  number.NewNamedUintParameter("base", 256, 32, 10),
		index: number.NewNamedIntParameter("index", 128, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPGATHERQD128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.base, v.index, v.mask}
}

func (v *VPGATHERQD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERQD128) Name() string {
	return "VPGATHERQD (128 bit) "
}

func (v *VPGATHERQD128) Description() string {
	return "Gather 2 u32 elements from base + (i64 index * 4); lanes selected by mask (MSB). Lower 2 indices are used; result merged into src."
}

func (v *VPGATHERQD128) Stub() string {
	return stubVpgatherqd128
}

func (v *VPGATHERQD128) Assembly() string {
	return assemblyVpgatherqd128
}

func (v *VPGATHERQD128) Run() {
	base := [8]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	index := [2]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpgatherqd128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPGATHERQD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
