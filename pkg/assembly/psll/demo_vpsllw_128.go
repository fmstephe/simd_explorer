package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllw_128.s
var assemblyVpsllw128 string

//go:embed stub_vpsllw_128.go
var stubVpsllw128 string

type VPSLLW128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLW128() *VPSLLW128 {
	return &VPSLLW128{
		vals:  number.NewNamedUintParameter("vals", 128, 16, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSLLW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLW128) Name() string {
	return "VPSLLW (128 bit) "
}

func (v *VPSLLW128) Description() string {
	return "Logical left shift of packed 16-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSLLW128) Stub() string {
	return stubVpsllw128
}

func (v *VPSLLW128) Assembly() string {
	return assemblyVpsllw128
}

func (v *VPSLLW128) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]uint16{}

	vpsllw128(&vals, &shift, &ret)

	log.Printf("VPSLLW128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
