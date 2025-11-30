package pextrw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpextrw_128_seven_idx.s
var assemblyVpextrw128Seven_idx string

//go:embed stub_vpextrw_128_seven_idx.go
var stubVpextrw128Seven_idx string

type VPEXTRW128SEVEN_IDX struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPEXTRW128SEVEN_IDX() *VPEXTRW128SEVEN_IDX {
	return &VPEXTRW128SEVEN_IDX{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 10),
	}
}

func (v *VPEXTRW128SEVEN_IDX) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPEXTRW128SEVEN_IDX) Output() *number.Parameter {
	return v.ret
}

func (v *VPEXTRW128SEVEN_IDX) Name() string {
	return "VPEXTRW (128 bit) seven_idx"
}

func (v *VPEXTRW128SEVEN_IDX) Description() string {
	return "Extract 16-bit word at index 7 from XMM and zero-extend to 32 bits (VEX)."
}

func (v *VPEXTRW128SEVEN_IDX) Stub() string {
	return stubVpextrw128Seven_idx
}

func (v *VPEXTRW128SEVEN_IDX) Assembly() string {
	return assemblyVpextrw128Seven_idx
}

func (v *VPEXTRW128SEVEN_IDX) Run() (output []byte) {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	var ret uint32

	vpextrw128Seven_idx(&vals, &ret)

	log.Printf("VPEXTRW128SEVEN_IDX input %v output %v", vals, ret)

	out := number.Uint32ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VPEXTRW128SEVEN_IDX) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
