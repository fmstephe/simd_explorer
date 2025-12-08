package psra

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsraw_128.s
var assemblyVpsraw128 string

//go:embed stub_vpsraw_128.go
var stubVpsraw128 string

type VPSRAW128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRAW128() *VPSRAW128 {
	return &VPSRAW128{
		vals:  number.NewNamedIntParameter("vals", 128, 16, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPSRAW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRAW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAW128) Name() string {
	return "VPSRAW (128 bit) "
}

func (v *VPSRAW128) Description() string {
	return "Arithmetical right shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRAW128) Stub() string {
	return stubVpsraw128
}

func (v *VPSRAW128) Assembly() string {
	return assemblyVpsraw128
}

func (v *VPSRAW128) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]int16{}

	vpsraw128(&vals, &shift, &ret)

	log.Printf("VPSRAW128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
