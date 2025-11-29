package vpgatherqd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpgatherqd_256.s
var assemblyVpgatherqd256 string

//go:embed stub_vpgatherqd_256.go
var stubVpgatherqd256 string

type VPGATHERQD256 struct {
	base  *number.Parameter
	index *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVPGATHERQD256() *VPGATHERQD256 {
	return &VPGATHERQD256{
		base:  number.NewNamedUintParameter("base", 512, 32, 10),
		index: number.NewNamedIntParameter("index", 256, 64, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPGATHERQD256) Inputs() []*number.Parameter {
	return []*number.Parameter{v.base, v.index, v.mask}
}

func (v *VPGATHERQD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPGATHERQD256) Name() string {
	return "VPGATHERQD (256 bit) "
}

func (v *VPGATHERQD256) Description() string {
	return "Gather 4 u32 elements from base + (i64 index * 4); lanes selected by mask (MSB). Lower 4 indices are used; result merged into src."
}

func (v *VPGATHERQD256) Stub() string {
	return stubVpgatherqd256
}

func (v *VPGATHERQD256) Assembly() string {
	return assemblyVpgatherqd256
}

func (v *VPGATHERQD256) Run(_ [][]byte) (output []byte) {
	base := [16]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	index := [4]uint64{}
	copy(index[:], number.ToUint64Slice(v.index.FlatData()))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpgatherqd256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPGATHERQD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
