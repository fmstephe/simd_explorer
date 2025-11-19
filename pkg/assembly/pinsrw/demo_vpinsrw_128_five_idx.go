package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpinsrw_128_five_idx.s
var assemblyVpinsrw128Five_idx string

//go:embed stub_vpinsrw_128_five_idx.go
var stubVpinsrw128Five_idx string

type VPINSRW128FIVE_IDX struct {
}

func (v *VPINSRW128FIVE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(16, 16, 10),
	}
}

func (v *VPINSRW128FIVE_IDX) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *VPINSRW128FIVE_IDX) Name() string {
	return "VPINSRW (128 bit) five_idx"
}

func (v *VPINSRW128FIVE_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 5; other lanes preserved (VEX)."
}

func (v *VPINSRW128FIVE_IDX) Stub() string {
	return stubVpinsrw128Five_idx
}

func (v *VPINSRW128FIVE_IDX) Assembly() string {
	return assemblyVpinsrw128Five_idx
}

func (v *VPINSRW128FIVE_IDX) Run(inputs [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(inputs[0]))
	scalar := number.ToUint16(inputs[1])

	ret := [8]uint16{}

	vpinsrw128Five_idx(&base, &scalar, &ret)

	log.Printf("VPINSRW128FIVE_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPINSRW128FIVE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
