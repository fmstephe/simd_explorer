package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pextrw_128_zero_idx.s
var assemblyPextrw128Zero_idx string

//go:embed stub_pextrw_128_zero_idx.go
var stubPextrw128Zero_idx string

type PEXTRW128ZERO_IDX struct {
}

func (v *PEXTRW128ZERO_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *PEXTRW128ZERO_IDX) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 10)
}

func (v *PEXTRW128ZERO_IDX) Name() string {
	return "PEXTRW (128 bit) zero_idx"
}

func (v *PEXTRW128ZERO_IDX) Description() string {
	return "Extract 16-bit word at index 0 from XMM and zero-extend to 32 bits."
}

func (v *PEXTRW128ZERO_IDX) Stub() string {
	return stubPextrw128Zero_idx
}

func (v *PEXTRW128ZERO_IDX) Assembly() string {
	return assemblyPextrw128Zero_idx
}

func (v *PEXTRW128ZERO_IDX) Run(inputs [][]byte) (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(inputs[0]))

	var ret uint32

	pextrw128Zero_idx(&vals, &ret)

	log.Printf("PEXTRW128ZERO_IDX input %v output %v", vals, ret)

	return number.Uint32ToBytes(ret)
}

func (v *PEXTRW128ZERO_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
