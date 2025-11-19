package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pinsrw_128_seven_idx.s
var assemblyPinsrw128Seven_idx string

//go:embed stub_pinsrw_128_seven_idx.go
var stubPinsrw128Seven_idx string

type PINSRW128SEVEN_IDX struct {
}

func (v *PINSRW128SEVEN_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(16, 16, 10),
	}
}

func (v *PINSRW128SEVEN_IDX) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PINSRW128SEVEN_IDX) Name() string {
	return "PINSRW (128 bit) seven_idx"
}

func (v *PINSRW128SEVEN_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 7; other lanes preserved."
}

func (v *PINSRW128SEVEN_IDX) Stub() string {
	return stubPinsrw128Seven_idx
}

func (v *PINSRW128SEVEN_IDX) Assembly() string {
	return assemblyPinsrw128Seven_idx
}

func (v *PINSRW128SEVEN_IDX) Run(inputs [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(inputs[0]))
	scalar := number.ToUint16(inputs[1])

	ret := [8]uint16{}

	pinsrw128Seven_idx(&base, &scalar, &ret)

	log.Printf("PINSRW128SEVEN_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PINSRW128SEVEN_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
