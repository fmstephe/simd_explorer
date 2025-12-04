package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubq_128.s
var assemblyVpsubq128 string

//go:embed stub_vpsubq_128.go
var stubVpsubq128 string

type VPSUBQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBQ128() *VPSUBQ128 {
	return &VPSUBQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPSUBQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBQ128) Name() string {
	return "VPSUBQ (128 bit) "
}

func (v *VPSUBQ128) Description() string {
	return "Subtract packed u64 quadwords (wrap-around)."
}

func (v *VPSUBQ128) Stub() string {
	return stubVpsubq128
}

func (v *VPSUBQ128) Assembly() string {
	return assemblyVpsubq128
}

func (v *VPSUBQ128) Run() {
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpsubq128(&vals1, &vals2, &ret)

	log.Printf("VPSUBQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSUBQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
