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
}

func (v *VPGATHERQD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10), // base memory (8 x u32)
		number.NewIntParameter(128, 64, 10),  // indices (i64; lower 2 used)
		number.NewUintParameter(128, 32, 16), // mask (MSB of each dword lane)
		number.NewUintParameter(128, 32, 10), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERQD128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10) // gathered vector
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

func (v *VPGATHERQD128) Run(inputs [][]byte) (output []byte) {
	base := [8]uint32{}
	copy(base[:], number.ToUint32Slice(inputs[0]))
	index := [2]uint64{}
	copy(index[:], number.ToUint64Slice(inputs[1]))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(inputs[2]))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(inputs[3]))

	vpgatherqd128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPGATHERQD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
