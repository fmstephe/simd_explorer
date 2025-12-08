package psra

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrad_128.s
var assemblyVpsrad128 string

//go:embed stub_vpsrad_128.go
var stubVpsrad128 string

type VPSRAD128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRAD128() *VPSRAD128 {
	return &VPSRAD128{
		vals:  number.NewNamedIntParameter("vals", 128, 32, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPSRAD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRAD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAD128) Name() string {
	return "VPSRAD (128 bit) "
}

func (v *VPSRAD128) Description() string {
	return "Arithmetical right shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRAD128) Stub() string {
	return stubVpsrad128
}

func (v *VPSRAD128) Assembly() string {
	return assemblyVpsrad128
}

func (v *VPSRAD128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [4]int32{}

	vpsrad128(&vals, &shift, &ret)

	log.Printf("VPSRAD128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
