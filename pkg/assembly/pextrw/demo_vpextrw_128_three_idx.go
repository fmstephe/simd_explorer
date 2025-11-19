package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpextrw_128_three_idx.s
var assemblyVpextrw128Three_idx string

//go:embed stub_vpextrw_128_three_idx.go
var stubVpextrw128Three_idx string

type VPEXTRW128THREE_IDX struct {
}

func (v *VPEXTRW128THREE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *VPEXTRW128THREE_IDX) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 10)
}

func (v *VPEXTRW128THREE_IDX) Name() string {
	return "VPEXTRW (128 bit) three_idx"
}

func (v *VPEXTRW128THREE_IDX) Description() string {
	return "Extract 16-bit word at index 3 from XMM and zero-extend to 32 bits (VEX)."
}

func (v *VPEXTRW128THREE_IDX) Stub() string {
	return stubVpextrw128Three_idx
}

func (v *VPEXTRW128THREE_IDX) Assembly() string {
	return assemblyVpextrw128Three_idx
}

func (v *VPEXTRW128THREE_IDX) Run(inputs [][]byte) (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(inputs[0]))

	var ret uint32

	vpextrw128Three_idx(&vals, &ret)

	log.Printf("VPEXTRW128THREE_IDX input %v output %v", vals, ret)

	return number.Uint32ToBytes(ret)
}

func (v *VPEXTRW128THREE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
