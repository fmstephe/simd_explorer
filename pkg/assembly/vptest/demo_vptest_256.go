package vptest

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vptest_256.s
var assemblyVptest256 string

//go:embed stub_vptest_256.go
var stubVptest256 string

type VPTEST256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPTEST256() *VPTEST256 {
	return &VPTEST256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 10),
	}
}

func (v *VPTEST256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPTEST256) Output() *number.Parameter {
	return v.ret
}

func (v *VPTEST256) Name() string {
	return "VPTEST (256 bit) "
}

func (v *VPTEST256) Description() string {
	return "Test integer vector bits. ZF=1 if (a & b) is all zeros; CF=1 if (~a & b) is all zeros. " +
		"Output encodes flags as bit0=ZF, bit1=CF."
}

func (v *VPTEST256) Stub() string {
	return stubVptest256
}

func (v *VPTEST256) Assembly() string {
	return assemblyVptest256
}

func (v *VPTEST256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], number.ToUint8Slice(v.vals1.FlatData()))
	vals2 := [32]uint8{}
	copy(vals2[:], number.ToUint8Slice(v.vals2.FlatData()))
	var ret uint32

	vptest256(&vals1, &vals2, &ret)

	log.Printf("VPTEST vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Uint32ToBytes(ret))
}

func (v *VPTEST256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
