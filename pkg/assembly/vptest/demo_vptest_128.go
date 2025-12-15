package vptest

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vptest_128.s
var assemblyVptest128 string

//go:embed stub_vptest_128.go
var stubVptest128 string

type VPTEST128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPTEST128() *VPTEST128 {
	return &VPTEST128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 10),
	}
}

func (v *VPTEST128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPTEST128) Output() *number.Parameter {
	return v.ret
}

func (v *VPTEST128) Name() string {
	return "VPTEST (128 bit) "
}

func (v *VPTEST128) Description() string {
	return "Test integer vector bits. ZF=1 if (a & b) is all zeros; CF=1 if (~a & b) is all zeros. " +
		"Output encodes flags as bit0=ZF, bit1=CF."
}

func (v *VPTEST128) Stub() string {
	return stubVptest128
}

func (v *VPTEST128) Assembly() string {
	return assemblyVptest128
}

func (v *VPTEST128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], number.ToUint8Slice(v.vals1.FlatData()))
	vals2 := [16]uint8{}
	copy(vals2[:], number.ToUint8Slice(v.vals2.FlatData()))
	var ret uint32

	vptest128(&vals1, &vals2, &ret)

	log.Printf("VPTEST vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(number.Uint32ToBytes(ret))
}

func (v *VPTEST128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
