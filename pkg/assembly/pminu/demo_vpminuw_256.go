package pminu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminuw_256.s
var assemblyVpminuw256 string

//go:embed stub_vpminuw_256.go
var stubVpminuw256 string

type VPMINUW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUW256() *VPMINUW256 {
	return &VPMINUW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPMINUW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINUW256) Name() string {
	return "VPMINUW (256 bit) "
}

func (v *VPMINUW256) Description() string {
	return "Unsigned minimum of packed 16-bit integers."
}

func (v *VPMINUW256) Stub() string {
	return stubVpminuw256
}

func (v *VPMINUW256) Assembly() string {
	return assemblyVpminuw256
}

func (v *VPMINUW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpminuw256(&vals1, &vals2, &ret)

	log.Printf("VPMINUW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINUW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
