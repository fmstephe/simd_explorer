package pinsrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpinsrw_128_one_idx.s
var assemblyVpinsrw128One_idx string

//go:embed stub_vpinsrw_128_one_idx.go
var stubVpinsrw128One_idx string

type VPINSRW128ONE_IDX struct {
	base   *number.Parameter
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPINSRW128ONE_IDX() *VPINSRW128ONE_IDX {
	return &VPINSRW128ONE_IDX{
		base:   number.NewNamedUintParameter("base", 128, 16, 10),
		scalar: number.NewNamedUintParameter("scalar", 16, 16, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPINSRW128ONE_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.scalar,
	}
}

func (v *VPINSRW128ONE_IDX) Output() *number.Parameter {
	return v.ret
}

func (v *VPINSRW128ONE_IDX) Name() string {
	return "VPINSRW (128 bit) one_idx"
}

func (v *VPINSRW128ONE_IDX) Description() string {
	return "Insert 16-bit word (scalar) into XMM at index 1; other lanes preserved (VEX)."
}

func (v *VPINSRW128ONE_IDX) Stub() string {
	return stubVpinsrw128One_idx
}

func (v *VPINSRW128ONE_IDX) Assembly() string {
	return assemblyVpinsrw128One_idx
}

func (v *VPINSRW128ONE_IDX) Run() {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(v.base.FlatData()))
	scalar := number.ToUint16(v.scalar.FlatData())

	ret := [8]uint16{}

	vpinsrw128One_idx(&base, &scalar, &ret)

	log.Printf("VPINSRW128ONE_IDX input base=%v scalar=%v output %v", base, scalar, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPINSRW128ONE_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
