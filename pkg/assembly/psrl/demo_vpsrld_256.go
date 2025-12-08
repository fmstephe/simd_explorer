package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrld_256.s
var assemblyVpsrld256 string

//go:embed stub_vpsrld_256.go
var stubVpsrld256 string

type VPSRLD256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLD256() *VPSRLD256 {
	return &VPSRLD256{
		vals:  number.NewNamedUintParameter("vals", 256, 32, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSRLD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLD256) Name() string {
	return "VPSRLD (256 bit) "
}

func (v *VPSRLD256) Description() string {
	return "Logical right shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRLD256) Stub() string {
	return stubVpsrld256
}

func (v *VPSRLD256) Assembly() string {
	return assemblyVpsrld256
}

func (v *VPSRLD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]uint32{}

	vpsrld256(&vals, &shift, &ret)

	log.Printf("VPSRLD256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
