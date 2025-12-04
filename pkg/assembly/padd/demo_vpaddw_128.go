package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddw_128.s
var assemblyVpaddw128 string

//go:embed stub_vpaddw_128.go
var stubVpaddw128 string

type VPADDW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDW128() *VPADDW128 {
	return &VPADDW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPADDW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDW128) Name() string {
	return "VPADDW (128 bit) "
}

func (v *VPADDW128) Description() string {
	return "Add packed u16 words (wrap-around)."
}

func (v *VPADDW128) Stub() string {
	return stubVpaddw128
}

func (v *VPADDW128) Assembly() string {
	return assemblyVpaddw128
}

func (v *VPADDW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpaddw128(&vals1, &vals2, &ret)

	log.Printf("VPADDW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
