package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslld_128.s
var assemblyVpslld128 string

//go:embed stub_vpslld_128.go
var stubVpslld128 string

type VPSLLD128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLD128() *VPSLLD128 {
	return &VPSLLD128{
		vals:  number.NewNamedUintParameter("vals", 128, 32, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSLLD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLD128) Name() string {
	return "VPSLLD (128 bit) "
}

func (v *VPSLLD128) Description() string {
	return "Logical left shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSLLD128) Stub() string {
	return stubVpslld128
}

func (v *VPSLLD128) Assembly() string {
	return assemblyVpslld128
}

func (v *VPSLLD128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [4]uint32{}

	vpslld128(&vals, &shift, &ret)

	log.Printf("VPSLLD128 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
