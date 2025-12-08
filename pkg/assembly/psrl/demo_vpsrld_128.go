package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrld_128.s
var assemblyVpsrld128 string

//go:embed stub_vpsrld_128.go
var stubVpsrld128 string

type VPSRLD128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLD128() *VPSRLD128 {
	return &VPSRLD128{
		vals:  number.NewNamedUintParameter("vals", 128, 32, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSRLD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLD128) Name() string {
	return "VPSRLD (128 bit) "
}

func (v *VPSRLD128) Description() string {
	return "Logical right shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRLD128) Stub() string {
	return stubVpsrld128
}

func (v *VPSRLD128) Assembly() string {
	return assemblyVpsrld128
}

func (v *VPSRLD128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [4]uint32{}

	vpsrld128(&vals, &shift, &ret)

	log.Printf("VPSRLD128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
