package pminu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminud_128.s
var assemblyVpminud128 string

//go:embed stub_vpminud_128.go
var stubVpminud128 string

type VPMINUD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUD128() *VPMINUD128 {
	return &VPMINUD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMINUD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINUD128) Name() string {
	return "VPMINUD (128 bit) "
}

func (v *VPMINUD128) Description() string {
	return "Unsigned minimum of packed 32-bit integers."
}

func (v *VPMINUD128) Stub() string {
	return stubVpminud128
}

func (v *VPMINUD128) Assembly() string {
	return assemblyVpminud128
}

func (v *VPMINUD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpminud128(&vals1, &vals2, &ret)

	log.Printf("VPMINUD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINUD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
