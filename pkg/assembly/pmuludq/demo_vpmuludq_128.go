package pmuludq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmuludq_128.s
var assemblyVpmuludq128 string

//go:embed stub_vpmuludq_128.go
var stubVpmuludq128 string

type VPMULUDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULUDQ128() *VPMULUDQ128 {
	return &VPMULUDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMULUDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULUDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULUDQ128) Name() string {
	return "VPMULUDQ (128 bit) "
}

func (v *VPMULUDQ128) Description() string {
	return "Multiply pairs of unsigned 32-bit integers to 64-bit results (even lanes)."
}

func (v *VPMULUDQ128) Stub() string {
	return stubVpmuludq128
}

func (v *VPMULUDQ128) Assembly() string {
	return assemblyVpmuludq128
}

func (v *VPMULUDQ128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpmuludq128(&vals1, &vals2, &ret)

	log.Printf("VPMULUDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULUDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
