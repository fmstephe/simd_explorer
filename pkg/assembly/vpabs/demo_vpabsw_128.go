package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsw_128.s
var assemblyVpabsw128 string

//go:embed stub_vpabsw_128.go
var stubVpabsw128 string

type VPABSW128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSW128() *VPABSW128 {
	return &VPABSW128{
		vals: number.NewNamedIntParameter("vals", 128, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPABSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSW128) Name() string {
	return "VPABSW (128 bit) "
}

func (v *VPABSW128) Description() string {
	return "Absolute value of packed signed 16-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int16."
}

func (v *VPABSW128) Stub() string {
	return stubVpabsw128
}

func (v *VPABSW128) Assembly() string {
	return assemblyVpabsw128
}

func (v *VPABSW128) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [8]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpabsw128(&vals, &ret)

	log.Printf("VPABSW vals %v ret %v", vals, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
