package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsd_128.s
var assemblyVpabsd128 string

//go:embed stub_vpabsd_128.go
var stubVpabsd128 string

type VPABSD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSD128() *VPABSD128 {
	return &VPABSD128{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPABSD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSD128) Name() string {
	return "VPABSD (128 bit) "
}

func (v *VPABSD128) Description() string {
	return "Absolute value of packed signed 32-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int32."
}

func (v *VPABSD128) Stub() string {
	return stubVpabsd128
}

func (v *VPABSD128) Assembly() string {
	return assemblyVpabsd128
}

func (v *VPABSD128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpabsd128(&vals, &ret)

	log.Printf("VPABSD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
