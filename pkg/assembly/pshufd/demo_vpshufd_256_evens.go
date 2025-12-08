package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_256_evens.s
var assemblyVpshufd256Evens string

//go:embed stub_vpshufd_256_evens.go
var stubVpshufd256Evens string

type VPSHUFD256EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD256EVENS() *VPSHUFD256EVENS {
	return &VPSHUFD256EVENS{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSHUFD256EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD256EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD256EVENS) Name() string {
	return "VPSHUFD EVENS (256 bit) "
}

func (v *VPSHUFD256EVENS) Description() string {
	return "Shuffle 32-bit integers within each 128-bit lane using imm8=0x88 (evens: [0,2,0,2])."
}

func (v *VPSHUFD256EVENS) Stub() string {
	return stubVpshufd256Evens
}

func (v *VPSHUFD256EVENS) Assembly() string {
	return assemblyVpshufd256Evens
}

func (v *VPSHUFD256EVENS) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [8]uint32{}

	vpshufd256Evens(&vals, &ret)

	log.Printf("VPSHUFD256EVENS vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD256EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
