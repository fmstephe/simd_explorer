package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsw_256.s
var assemblyVpabsw256 string

//go:embed stub_vpabsw_256.go
var stubVpabsw256 string

type VPABSW256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSW256() *VPABSW256 {
	return &VPABSW256{
		vals: number.NewNamedIntParameter("vals", 256, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPABSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSW256) Name() string {
	return "VPABSW (256 bit) "
}

func (v *VPABSW256) Description() string {
	return "Absolute value of packed signed 16-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int16; operates per 128-bit lane."
}

func (v *VPABSW256) Stub() string {
	return stubVpabsw256
}

func (v *VPABSW256) Assembly() string {
	return assemblyVpabsw256
}

func (v *VPABSW256) Run() {
	vals := [16]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [16]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpabsw256(&vals, &ret)

	log.Printf("VPABSW vals %v ret %v", vals, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
