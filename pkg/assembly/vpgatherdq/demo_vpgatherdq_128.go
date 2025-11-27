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
}

func (v *VPGATHERDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 64, 10), // base memory (8 x u32)
		number.NewIntParameter(128, 32, 10),  // indices (i64; lower 2 used)
		number.NewUintParameter(128, 64, 16), // mask (MSB per dword lane)
		number.NewUintParameter(128, 64, 10), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERDQ128) Output() *number.Parameter {
	return number.NewUintParameter(128, 64, 10) // gathered vector
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

func (v *VPGATHERDQ128) Run(inputs [][]byte) (output []byte) {
	base := [4]uint64{}
	copy(base[:], number.ToUint64Slice(inputs[0]))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1]))
	mask := [2]uint64{}
	copy(mask[:], number.ToUint64Slice(inputs[2]))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(inputs[3]))

	vpgatherdq128(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDQ128 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint64SliceToBytes(ret[:])
}

func (v *VPGATHERDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
