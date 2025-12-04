package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubw_256.s
var assemblyVpsubw256 string

//go:embed stub_vpsubw_256.go
var stubVpsubw256 string

type VPSUBW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBW256() *VPSUBW256 {
	return &VPSUBW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSUBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBW256) Name() string {
	return "VPSUBW (256 bit) "
}

func (v *VPSUBW256) Description() string {
	return "Subtract packed u16 words (wrap-around)."
}

func (v *VPSUBW256) Stub() string {
	return stubVpsubw256
}

func (v *VPSUBW256) Assembly() string {
	return assemblyVpsubw256
}

func (v *VPSUBW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpsubw256(&vals1, &vals2, &ret)

	log.Printf("VPSUBW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSUBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
