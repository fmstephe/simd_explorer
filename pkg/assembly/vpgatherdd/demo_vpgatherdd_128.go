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
}

func (v *VPGATHERDD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 16), // base memory (8 x u32)
		number.NewIntParameter(128, 32, 10),  // indices (i32; lower 4 used)
		number.NewUintParameter(128, 32, 16), // mask (MSB of each dword lane)
		number.NewUintParameter(128, 32, 16), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERDD128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16) // gathered vector
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

func (v *VPGATHERDD128) Run(inputs [][]byte) (output []byte) {
	base := [8]uint32{}
	copy(base[:], number.ToUint32Slice(inputs[0]))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1])) // UI uses base-10; we read bits
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(inputs[2])) // hex mask; MSB per dword
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(inputs[3])) // src/dst seed

	vpgatherdd128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDD128 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPGATHERDD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
