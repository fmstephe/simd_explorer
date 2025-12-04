package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddq_128.s
var assemblyVpaddq128 string

//go:embed stub_vpaddq_128.go
var stubVpaddq128 string

type VPADDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDQ128() *VPADDQ128 {
	return &VPADDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPADDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDQ128) Name() string {
	return "VPADDQ (128 bit) "
}

func (v *VPADDQ128) Description() string {
	return "Add packed u64 qwords (wrap-around)."
}

func (v *VPADDQ128) Stub() string {
	return stubVpaddq128
}

func (v *VPADDQ128) Assembly() string {
	return assemblyVpaddq128
}

func (v *VPADDQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpaddq128(&vals1, &vals2, &ret)

	log.Printf("VPADDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
