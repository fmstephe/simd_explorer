package psra

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsraw_256.s
var assemblyVpsraw256 string

//go:embed stub_vpsraw_256.go
var stubVpsraw256 string

type VPSRAW256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRAW256() *VPSRAW256 {
	return &VPSRAW256{
		vals:  number.NewNamedIntParameter("vals", 256, 16, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPSRAW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRAW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAW256) Name() string {
	return "VPSRAW (256 bit) "
}

func (v *VPSRAW256) Description() string {
	return "Arithmetical right shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRAW256) Stub() string {
	return stubVpsraw256
}

func (v *VPSRAW256) Assembly() string {
	return assemblyVpsraw256
}

func (v *VPSRAW256) Run() {
	vals := [16]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [16]int16{}

	vpsraw256(&vals, &shift, &ret)

	log.Printf("VPSRAW256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
