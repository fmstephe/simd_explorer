package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_256_evens.s
var assemblyVpshuflw256Evens string

//go:embed stub_vpshuflw_256_evens.go
var stubVpshuflw256Evens string

type VPSHUFLW256EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW256EVENS() *VPSHUFLW256EVENS {
	return &VPSHUFLW256EVENS{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFLW256EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW256EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW256EVENS) Name() string {
	return "VPSHUFLW (256 bit) evens"
}

func (v *VPSHUFLW256EVENS) Description() string {
	return "Shuffle low words per 128-bit lane using imm8=0x88 (evens: [w4,w6,w4,w6]); high words unchanged."
}

func (v *VPSHUFLW256EVENS) Stub() string {
	return stubVpshuflw256Evens
}

func (v *VPSHUFLW256EVENS) Assembly() string {
	return assemblyVpshuflw256Evens
}

func (v *VPSHUFLW256EVENS) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshuflw256Evens(&vals, &ret)

	log.Printf("VPSHUFLW256EVENS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW256EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
