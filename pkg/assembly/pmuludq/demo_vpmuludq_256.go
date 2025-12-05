package pmuludq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmuludq_256.s
var assemblyVpmuludq256 string

//go:embed stub_vpmuludq_256.go
var stubVpmuludq256 string

type VPMULUDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULUDQ256() *VPMULUDQ256 {
	return &VPMULUDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPMULUDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULUDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULUDQ256) Name() string {
	return "VPMULUDQ (256 bit) "
}

func (v *VPMULUDQ256) Description() string {
	return "Multiply pairs of unsigned 32-bit integers to 64-bit results (even lanes)."
}

func (v *VPMULUDQ256) Stub() string {
	return stubVpmuludq256
}

func (v *VPMULUDQ256) Assembly() string {
	return assemblyVpmuludq256
}

func (v *VPMULUDQ256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpmuludq256(&vals1, &vals2, &ret)

	log.Printf("VPMULUDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULUDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
