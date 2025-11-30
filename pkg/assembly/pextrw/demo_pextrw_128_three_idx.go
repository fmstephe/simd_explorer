package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pextrw_128_three_idx.s
var assemblyPextrw128Three_idx string

//go:embed stub_pextrw_128_three_idx.go
var stubPextrw128Three_idx string

type PEXTRW128THREE_IDX struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewPEXTRW128THREE_IDX() *PEXTRW128THREE_IDX {
	return &PEXTRW128THREE_IDX{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 10),
	}
}

func (v *PEXTRW128THREE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *PEXTRW128THREE_IDX) Output() *number.Parameter {
	return v.ret
}

func (v *PEXTRW128THREE_IDX) Name() string {
	return "PEXTRW (128 bit) three_idx"
}

func (v *PEXTRW128THREE_IDX) Description() string {
	return "Extract 16-bit word at index 3 from XMM and zero-extend to 32 bits."
}

func (v *PEXTRW128THREE_IDX) Stub() string {
	return stubPextrw128Three_idx
}

func (v *PEXTRW128THREE_IDX) Assembly() string {
	return assemblyPextrw128Three_idx
}

func (v *PEXTRW128THREE_IDX) Run() (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	var ret uint32

	pextrw128Three_idx(&vals, &ret)

	log.Printf("PEXTRW128THREE_IDX input %v output %v", vals, ret)

	out := number.Uint32ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *PEXTRW128THREE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
