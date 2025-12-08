package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslld_256.s
var assemblyVpslld256 string

//go:embed stub_vpslld_256.go
var stubVpslld256 string

type VPSLLD256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLD256() *VPSLLD256 {
	return &VPSLLD256{
		vals:  number.NewNamedUintParameter("vals", 256, 32, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSLLD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLD256) Name() string {
	return "VPSLLD (256 bit) "
}

func (v *VPSLLD256) Description() string {
	return "Logical left shift of packed 32-bit integers by per 128-bit-lane 64-bit counts."
}

func (v *VPSLLD256) Stub() string {
	return stubVpslld256
}

func (v *VPSLLD256) Assembly() string {
	return assemblyVpslld256
}

func (v *VPSLLD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [8]uint32{}

	vpslld256(&vals, &shift, &ret)

	log.Printf("VPSLLD256 vals %v shift(lanes) %v ret %v", vals, shift, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
