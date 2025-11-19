package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pextrw_128_four_idx.s
var assemblyPextrw128Four_idx string

//go:embed stub_pextrw_128_four_idx.go
var stubPextrw128Four_idx string

type PEXTRW128FOUR_IDX struct {
}

func (v *PEXTRW128FOUR_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *PEXTRW128FOUR_IDX) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 10)
}

func (v *PEXTRW128FOUR_IDX) Name() string {
	return "PEXTRW (128 bit) four_idx"
}

func (v *PEXTRW128FOUR_IDX) Description() string {
	return "Extract 16-bit word at index 4 from XMM and zero-extend to 32 bits."
}

func (v *PEXTRW128FOUR_IDX) Stub() string {
	return stubPextrw128Four_idx
}

func (v *PEXTRW128FOUR_IDX) Assembly() string {
	return assemblyPextrw128Four_idx
}

func (v *PEXTRW128FOUR_IDX) Run(inputs [][]byte) (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(inputs[0]))

	var ret uint32

	pextrw128Four_idx(&vals, &ret)

	log.Printf("PEXTRW128FOUR_IDX input %v output %v", vals, ret)

	return number.Uint32ToBytes(ret)
}

func (v *PEXTRW128FOUR_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
