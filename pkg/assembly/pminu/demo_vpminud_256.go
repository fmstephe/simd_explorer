package pminu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminud_256.s
var assemblyVpminud256 string

//go:embed stub_vpminud_256.go
var stubVpminud256 string

type VPMINUD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUD256() *VPMINUD256 {
	return &VPMINUD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMINUD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINUD256) Name() string {
	return "VPMINUD (256 bit) "
}

func (v *VPMINUD256) Description() string {
	return "Unsigned minimum of packed 32-bit integers."
}

func (v *VPMINUD256) Stub() string {
	return stubVpminud256
}

func (v *VPMINUD256) Assembly() string {
	return assemblyVpminud256
}

func (v *VPMINUD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpminud256(&vals1, &vals2, &ret)

	log.Printf("VPMINUD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINUD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
