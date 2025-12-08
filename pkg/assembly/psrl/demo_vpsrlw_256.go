package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlw_256.s
var assemblyVpsrlw256 string

//go:embed stub_vpsrlw_256.go
var stubVpsrlw256 string

type VPSRLW256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLW256() *VPSRLW256 {
	return &VPSRLW256{
		vals:  number.NewNamedUintParameter("vals", 256, 16, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSRLW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLW256) Name() string {
	return "VPSRLW (256 bit) "
}

func (v *VPSRLW256) Description() string {
	return "Logical right shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSRLW256) Stub() string {
	return stubVpsrlw256
}

func (v *VPSRLW256) Assembly() string {
	return assemblyVpsrlw256
}

func (v *VPSRLW256) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [16]uint16{}

	vpsrlw256(&vals, &shift, &ret)

	log.Printf("VPSRLW256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
