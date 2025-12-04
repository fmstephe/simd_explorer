package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubw_128.s
var assemblyVpsubw128 string

//go:embed stub_vpsubw_128.go
var stubVpsubw128 string

type VPSUBW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBW128() *VPSUBW128 {
	return &VPSUBW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSUBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBW128) Name() string {
	return "VPSUBW (128 bit) "
}

func (v *VPSUBW128) Description() string {
	return "Subtract packed u16 words (wrap-around)."
}

func (v *VPSUBW128) Stub() string {
	return stubVpsubw128
}

func (v *VPSUBW128) Assembly() string {
	return assemblyVpsubw128
}

func (v *VPSUBW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpsubw128(&vals1, &vals2, &ret)

	log.Printf("VPSUBW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSUBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
