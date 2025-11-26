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
}

func (v *VPGATHERQD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(512, 32, 10), // base memory (16 x u32)
		number.NewIntParameter(256, 64, 10),  // indices (i64; lower 4 used)
		number.NewUintParameter(128, 32, 16), // mask (MSB of each dword lane)
		number.NewUintParameter(128, 32, 10), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERQD256) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10) // gathered vector
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

func (v *VPGATHERQD256) Run(inputs [][]byte) (output []byte) {
	base := [16]uint32{}
	copy(base[:], number.ToUint32Slice(inputs[0]))
	index := [4]uint64{}
	copy(index[:], number.ToUint64Slice(inputs[1]))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(inputs[2]))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(inputs[3]))

	vpgatherqd256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERQD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPGATHERQD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
