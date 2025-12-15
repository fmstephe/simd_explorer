package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsd_256.s
var assemblyVpabsd256 string

//go:embed stub_vpabsd_256.go
var stubVpabsd256 string

type VPABSD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSD256() *VPABSD256 {
	return &VPABSD256{
		vals: number.NewNamedIntParameter("vals", 256, 32, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPABSD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSD256) Name() string {
	return "VPABSD (256 bit) "
}

func (v *VPABSD256) Description() string {
	return "Absolute value of packed signed 32-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int32; operates per 128-bit lane."
}

func (v *VPABSD256) Stub() string {
	return stubVpabsd256
}

func (v *VPABSD256) Assembly() string {
	return assemblyVpabsd256
}

func (v *VPABSD256) Run() {
	vals := [8]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpabsd256(&vals, &ret)

	log.Printf("VPABSD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
