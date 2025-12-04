package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubq_256.s
var assemblyVpsubq256 string

//go:embed stub_vpsubq_256.go
var stubVpsubq256 string

type VPSUBQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBQ256() *VPSUBQ256 {
	return &VPSUBQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPSUBQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBQ256) Name() string {
	return "VPSUBQ (256 bit) "
}

func (v *VPSUBQ256) Description() string {
	return "Subtract packed u64 quadwords (wrap-around)."
}

func (v *VPSUBQ256) Stub() string {
	return stubVpsubq256
}

func (v *VPSUBQ256) Assembly() string {
	return assemblyVpsubq256
}

func (v *VPSUBQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpsubq256(&vals1, &vals2, &ret)

	log.Printf("VPSUBQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSUBQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
