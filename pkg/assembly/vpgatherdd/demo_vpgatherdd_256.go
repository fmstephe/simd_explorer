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
}

func (v *VPGATHERDD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(512, 32, 16), // base memory (32 x u32)
		number.NewIntParameter(256, 32, 10),  // indices (i32; lower 8 used)
		number.NewUintParameter(256, 32, 16), // mask (MSB of each dword lane)
		number.NewUintParameter(256, 32, 16), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERDD256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16) // gathered vector
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
	copy(base[:], number.ToUint32Slice(inputs[0]))
	index := [8]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1]))
	mask := [8]uint32{}
	copy(mask[:], number.ToUint32Slice(inputs[2]))
	ret := [8]uint32{}
	copy(ret[:], number.ToUint32Slice(inputs[3]))

	vpgatherdd256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDD256 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPGATHERDD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
