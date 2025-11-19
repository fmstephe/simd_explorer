package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pinsrw_128_zero_idx.s
var assemblyPinsrw128Zero_idx string

//go:embed stub_pinsrw_128_zero_idx.go
var stubPinsrw128Zero_idx string

type PINSRW128ZERO_IDX struct {
}

func (v *PINSRW128ZERO_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(16, 16, 10),
	}
}

func (v *PINSRW128ZERO_IDX) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PINSRW128ZERO_IDX) Name() string {
	return "PINSRW (128 bit) zero_idx"
}

func (v *PINSRW128ZERO_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 0; other lanes preserved."
}

func (v *PINSRW128ZERO_IDX) Stub() string {
	return stubPinsrw128Zero_idx
}

func (v *PINSRW128ZERO_IDX) Assembly() string {
	return assemblyPinsrw128Zero_idx
}

func (v *PINSRW128ZERO_IDX) Run(inputs [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(inputs[0]))
	scalar := number.ToUint16(inputs[1])

	ret := [8]uint16{}

	pinsrw128Zero_idx(&base, &scalar, &ret)

	log.Printf("PINSRW128ZERO_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PINSRW128ZERO_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
