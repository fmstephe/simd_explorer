package vpabs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabsb_256.s
var assemblyVpabsb256 string

//go:embed stub_vpabsb_256.go
var stubVpabsb256 string

type VPABSB256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPABSB256() *VPABSB256 {
	return &VPABSB256{
		vals: number.NewNamedIntParameter("vals", 256, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 8, 10),
	}
}

func (v *VPABSB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPABSB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPABSB256) Name() string {
	return "VPABSB (256 bit) "
}

func (v *VPABSB256) Description() string {
	return "Absolute value of packed signed 8-bit integers. " +
		"Each lane computes abs(a) with signed saturation into int8; operates per 128-bit lane."
}

func (v *VPABSB256) Stub() string {
	return stubVpabsb256
}

func (v *VPABSB256) Assembly() string {
	return assemblyVpabsb256
}

func (v *VPABSB256) Run() {
	vals := [32]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [32]int8{}
	copy(ret[:], number.ToInt8Slice(v.ret.FlatData()))

	vpabsb256(&vals, &ret)

	log.Printf("VPABSB vals %v ret %v", vals, ret)

	retBytes := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPABSB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
