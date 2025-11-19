package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pinsrw_128_three_idx.s
var assemblyPinsrw128Three_idx string

//go:embed stub_pinsrw_128_three_idx.go
var stubPinsrw128Three_idx string

type PINSRW128THREE_IDX struct {
}

func (v *PINSRW128THREE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(16, 16, 10),
	}
}

func (v *PINSRW128THREE_IDX) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PINSRW128THREE_IDX) Name() string {
	return "PINSRW (128 bit) three_idx"
}

func (v *PINSRW128THREE_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 3; other lanes preserved."
}

func (v *PINSRW128THREE_IDX) Stub() string {
	return stubPinsrw128Three_idx
}

func (v *PINSRW128THREE_IDX) Assembly() string {
	return assemblyPinsrw128Three_idx
}

func (v *PINSRW128THREE_IDX) Run(inputs [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(inputs[0]))
	scalar := number.ToUint16(inputs[1])

	ret := [8]uint16{}

	pinsrw128Three_idx(&base, &scalar, &ret)

	log.Printf("PINSRW128THREE_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PINSRW128THREE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
