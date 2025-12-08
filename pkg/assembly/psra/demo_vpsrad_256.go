package psra

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrad_256.s
var assemblyVpsrad256 string

//go:embed stub_vpsrad_256.go
var stubVpsrad256 string

type VPSRAD256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRAD256() *VPSRAD256 {
	return &VPSRAD256{
		vals:  number.NewNamedIntParameter("vals", 256, 32, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPSRAD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRAD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAD256) Name() string {
	return "VPSRAD (256 bit) "
}

func (v *VPSRAD256) Description() string {
	return "Arithmetical right shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRAD256) Stub() string {
	return stubVpsrad256
}

func (v *VPSRAD256) Assembly() string {
	return assemblyVpsrad256
}

func (v *VPSRAD256) Run() {
	vals := [8]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]int32{}

	vpsrad256(&vals, &shift, &ret)

	log.Printf("VPSRAD256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
