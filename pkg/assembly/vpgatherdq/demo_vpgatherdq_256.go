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
}

func (v *VPGATHERDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(512, 64, 10), // base memory (16 x u32)
		number.NewIntParameter(128, 32, 10),  // indices
		number.NewUintParameter(256, 64, 16), // mask (MSB per dword lane)
		number.NewUintParameter(256, 64, 10), // src/dst (merge for masked-off lanes)
	}
}

func (v *VPGATHERDQ256) Output() *number.Parameter {
	return number.NewUintParameter(256, 64, 10) // gathered vector
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

func (v *VPGATHERDQ256) Run(inputs [][]byte) (output []byte) {
	base := [8]uint64{}
	copy(base[:], number.ToUint64Slice(inputs[0]))
	index := [4]uint32{}
	copy(index[:], number.ToUint32Slice(inputs[1]))
	mask := [4]uint64{}
	copy(mask[:], number.ToUint64Slice(inputs[2]))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(inputs[3]))

	vpgatherdq256(&base, &index, &mask, &ret)

	log.Printf("VPGATHERDQ256 base %v index %v mask %v ret %v", base, index, mask, ret)

	return number.Uint64SliceToBytes(ret[:])
}

func (v *VPGATHERDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
