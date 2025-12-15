package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsb_128.s
var assemblyVpabsb128 string

//go:embed stub_vpabsb_128.go
var stubVpabsb128 string

type VPABSB128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSB128() *VPABSB128 {
	return &VPABSB128{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 8, 10),
	}
}

func (v *VPABSB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSB128) Name() string {
	return "VPABSB (128 bit) "
}

func (v *VPABSB128) Description() string {
	return "Absolute value of packed signed 8-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int8."
}

func (v *VPABSB128) Stub() string {
	return stubVpabsb128
}

func (v *VPABSB128) Assembly() string {
	return assemblyVpabsb128
}

func (v *VPABSB128) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [16]int8{}
	copy(ret[:], number.ToInt8Slice(v.ret.FlatData()))

	vpabsb128(&vals, &ret)

	log.Printf("VPABSB vals %v ret %v", vals, ret)

	retBytes := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
