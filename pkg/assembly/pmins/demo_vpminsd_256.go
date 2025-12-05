package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsd_256.s
var assemblyVpminsd256 string

//go:embed stub_vpminsd_256.go
var stubVpminsd256 string

type VPMINSD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSD256() *VPMINSD256 {
	return &VPMINSD256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMINSD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSD256) Name() string {
	return "VPMINSD (256 bit) "
}

func (v *VPMINSD256) Description() string {
	return "Signed minimum of packed 32-bit integers."
}

func (v *VPMINSD256) Stub() string {
	return stubVpminsd256
}

func (v *VPMINSD256) Assembly() string {
	return assemblyVpminsd256
}

func (v *VPMINSD256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [8]int32{}

	vpminsd256(&vals1, &vals2, &ret)

	log.Printf("VPMINSD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINSD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
