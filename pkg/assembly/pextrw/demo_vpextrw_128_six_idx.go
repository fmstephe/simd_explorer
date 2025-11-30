package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpextrw_128_six_idx.s
var assemblyVpextrw128Six_idx string

//go:embed stub_vpextrw_128_six_idx.go
var stubVpextrw128Six_idx string

type VPEXTRW128SIX_IDX struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPEXTRW128SIX_IDX() *VPEXTRW128SIX_IDX {
	return &VPEXTRW128SIX_IDX{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 10),
	}
}

func (v *VPEXTRW128SIX_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPEXTRW128SIX_IDX) Output() *number.Parameter {
	return v.ret
}

func (v *VPEXTRW128SIX_IDX) Name() string {
	return "VPEXTRW (128 bit) six_idx"
}

func (v *VPEXTRW128SIX_IDX) Description() string {
	return "Extract 16-bit word at index 6 from XMM and zero-extend to 32 bits (VEX)."
}

func (v *VPEXTRW128SIX_IDX) Stub() string {
	return stubVpextrw128Six_idx
}

func (v *VPEXTRW128SIX_IDX) Assembly() string {
	return assemblyVpextrw128Six_idx
}

func (v *VPEXTRW128SIX_IDX) Run() (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	var ret uint32

	vpextrw128Six_idx(&vals, &ret)

	log.Printf("VPEXTRW128SIX_IDX input %v output %v", vals, ret)

	out := number.Uint32ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VPEXTRW128SIX_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
