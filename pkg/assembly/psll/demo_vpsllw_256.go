package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllw_256.s
var assemblyVpsllw256 string

//go:embed stub_vpsllw_256.go
var stubVpsllw256 string

type VPSLLW256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLW256() *VPSLLW256 {
	return &VPSLLW256{
		vals:  number.NewNamedUintParameter("vals", 256, 16, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSLLW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLW256) Name() string {
	return "VPSLLW (256 bit) "
}

func (v *VPSLLW256) Description() string {
	return "Logical left shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSLLW256) Stub() string {
	return stubVpsllw256
}

func (v *VPSLLW256) Assembly() string {
	return assemblyVpsllw256
}

func (v *VPSLLW256) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [16]uint16{}

	vpsllw256(&vals, &shift, &ret)

	log.Printf("VPSLLW256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
