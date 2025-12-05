package pminu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminuw_128.s
var assemblyVpminuw128 string

//go:embed stub_vpminuw_128.go
var stubVpminuw128 string

type VPMINUW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUW128() *VPMINUW128 {
	return &VPMINUW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPMINUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINUW128) Name() string {
	return "VPMINUW (128 bit) "
}

func (v *VPMINUW128) Description() string {
	return "Unsigned minimum of packed 16-bit integers."
}

func (v *VPMINUW128) Stub() string {
	return stubVpminuw128
}

func (v *VPMINUW128) Assembly() string {
	return assemblyVpminuw128
}

func (v *VPMINUW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpminuw128(&vals1, &vals2, &ret)

	log.Printf("VPMINUW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
