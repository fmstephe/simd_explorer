package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_128_evens.s
var assemblyVpshufd128Evens string

//go:embed stub_vpshufd_128_evens.go
var stubVpshufd128Evens string

type VPSHUFD128EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD128EVENS() *VPSHUFD128EVENS {
	return &VPSHUFD128EVENS{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSHUFD128EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD128EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD128EVENS) Name() string {
	return "VPSHUFD EVENS (128 bit) "
}

func (v *VPSHUFD128EVENS) Description() string {
	return "Shuffle 32-bit integers within a 128-bit lane using imm8=0x88 (evens: [0,2,0,2])."
}

func (v *VPSHUFD128EVENS) Stub() string {
	return stubVpshufd128Evens
}

func (v *VPSHUFD128EVENS) Assembly() string {
	return assemblyVpshufd128Evens
}

func (v *VPSHUFD128EVENS) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vpshufd128Evens(&vals, &ret)

	log.Printf("VPSHUFD128EVENS vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD128EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
