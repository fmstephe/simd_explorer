package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddw_256.s
var assemblyVpaddw256 string

//go:embed stub_vpaddw_256.go
var stubVpaddw256 string

type VPADDW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDW256() *VPADDW256 {
	return &VPADDW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPADDW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDW256) Name() string {
	return "VPADDW (256 bit) "
}

func (v *VPADDW256) Description() string {
	return "Add packed u16 words (wrap-around)."
}

func (v *VPADDW256) Stub() string {
	return stubVpaddw256
}

func (v *VPADDW256) Assembly() string {
	return assemblyVpaddw256
}

func (v *VPADDW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpaddw256(&vals1, &vals2, &ret)

	log.Printf("VPADDW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
