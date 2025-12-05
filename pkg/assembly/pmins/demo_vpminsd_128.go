package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsd_128.s
var assemblyVpminsd128 string

//go:embed stub_vpminsd_128.go
var stubVpminsd128 string

type VPMINSD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSD128() *VPMINSD128 {
	return &VPMINSD128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMINSD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSD128) Name() string {
	return "VPMINSD (128 bit) "
}

func (v *VPMINSD128) Description() string {
	return "Signed minimum of packed 32-bit integers."
}

func (v *VPMINSD128) Stub() string {
	return stubVpminsd128
}

func (v *VPMINSD128) Assembly() string {
	return assemblyVpminsd128
}

func (v *VPMINSD128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [4]int32{}

	vpminsd128(&vals1, &vals2, &ret)

	log.Printf("VPMINSD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINSD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
