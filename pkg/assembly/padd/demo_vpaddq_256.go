package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddq_256.s
var assemblyVpaddq256 string

//go:embed stub_vpaddq_256.go
var stubVpaddq256 string

type VPADDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDQ256() *VPADDQ256 {
	return &VPADDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPADDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDQ256) Name() string {
	return "VPADDQ (256 bit) "
}

func (v *VPADDQ256) Description() string {
	return "Add packed u64 qwords (wrap-around)."
}

func (v *VPADDQ256) Stub() string {
	return stubVpaddq256
}

func (v *VPADDQ256) Assembly() string {
	return assemblyVpaddq256
}

func (v *VPADDQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpaddq256(&vals1, &vals2, &ret)

	log.Printf("VPADDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
