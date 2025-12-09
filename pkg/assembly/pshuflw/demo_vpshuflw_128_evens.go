package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_128_evens.s
var assemblyVpshuflw128Evens string

//go:embed stub_vpshuflw_128_evens.go
var stubVpshuflw128Evens string

type VPSHUFLW128EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW128EVENS() *VPSHUFLW128EVENS {
	return &VPSHUFLW128EVENS{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFLW128EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW128EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW128EVENS) Name() string {
	return "VPSHUFLW (128 bit) evens"
}

func (v *VPSHUFLW128EVENS) Description() string {
	return "Shuffle low words in 128-bit lane using imm8=0x88 (evens: [w4,w6,w4,w6]); high words unchanged."
}

func (v *VPSHUFLW128EVENS) Stub() string {
	return stubVpshuflw128Evens
}

func (v *VPSHUFLW128EVENS) Assembly() string {
	return assemblyVpshuflw128Evens
}

func (v *VPSHUFLW128EVENS) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshuflw128Evens(&vals, &ret)

	log.Printf("VPSHUFLW128EVENS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW128EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
