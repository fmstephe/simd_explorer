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
	base   *number.Parameter
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPINSRW128FIVE_IDX() *VPINSRW128FIVE_IDX {
	return &VPINSRW128FIVE_IDX{
		base:   number.NewNamedUintParameter("base", 128, 16, 10),
		scalar: number.NewNamedUintParameter("scalar", 16, 16, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPINSRW128FIVE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.scalar,
	}
}

func (v *VPINSRW128FIVE_IDX) Output() *number.Parameter {
	return v.ret
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

func (v *VPINSRW128FIVE_IDX) Run(_ [][]byte) (output []byte) {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(v.base.FlatData()))
	scalar := number.ToUint16(v.scalar.FlatData())

	ret := [8]uint16{}

	vpinsrw128Five_idx(&base, &scalar, &ret)

	log.Printf("VPINSRW128FIVE_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPINSRW128FIVE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
