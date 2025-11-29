package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pinsrw_128_four_idx.s
var assemblyPinsrw128Four_idx string

//go:embed stub_pinsrw_128_four_idx.go
var stubPinsrw128Four_idx string

type PINSRW128FOUR_IDX struct {
	base   *number.Parameter
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewPINSRW128FOUR_IDX() *PINSRW128FOUR_IDX {
	return &PINSRW128FOUR_IDX{
		base:   number.NewNamedUintParameter("base", 128, 16, 10),
		scalar: number.NewNamedUintParameter("scalar", 16, 16, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *PINSRW128FOUR_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.scalar,
	}
}

func (v *PINSRW128FOUR_IDX) Output() *number.Parameter {
	return v.ret
}

func (v *PINSRW128FOUR_IDX) Name() string {
	return "PINSRW (128 bit) four_idx"
}

func (v *PINSRW128FOUR_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 4; other lanes preserved."
}

func (v *PINSRW128FOUR_IDX) Stub() string {
	return stubPinsrw128Four_idx
}

func (v *PINSRW128FOUR_IDX) Assembly() string {
	return assemblyPinsrw128Four_idx
}

func (v *PINSRW128FOUR_IDX) Run(_ [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(v.base.FlatData()))
	scalar := number.ToUint16(v.scalar.FlatData())

	ret := [8]uint16{}

	pinsrw128Four_idx(&base, &scalar, &ret)

	log.Printf("PINSRW128FOUR_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *PINSRW128FOUR_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
