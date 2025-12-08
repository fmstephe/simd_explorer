package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlw_128.s
var assemblyVpsrlw128 string

//go:embed stub_vpsrlw_128.go
var stubVpsrlw128 string

type VPSRLW128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLW128() *VPSRLW128 {
	return &VPSRLW128{
		vals:  number.NewNamedUintParameter("vals", 128, 16, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSRLW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLW128) Name() string {
	return "VPSRLW (128 bit) "
}

func (v *VPSRLW128) Description() string {
	return "Logical right shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRLW128) Stub() string {
	return stubVpsrlw128
}

func (v *VPSRLW128) Assembly() string {
	return assemblyVpsrlw128
}

func (v *VPSRLW128) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]uint16{}

	vpsrlw128(&vals, &shift, &ret)

	log.Printf("VPSRLW128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
