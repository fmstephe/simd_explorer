package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckldq_256.s
var assemblyVpunpckldq256 string

//go:embed stub_vpunpckldq_256.go
var stubVpunpckldq256 string

type VPUNPCKLDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLDQ256() *VPUNPCKLDQ256 {
	return &VPUNPCKLDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPUNPCKLDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLDQ256) Name() string {
	return "VPUNPCKLDQ (256 bit) "
}

func (v *VPUNPCKLDQ256) Description() string {
	return "Unpack and interleave low-order doublewords from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKLDQ256) Stub() string {
	return stubVpunpckldq256
}

func (v *VPUNPCKLDQ256) Assembly() string {
	return assemblyVpunpckldq256
}

func (v *VPUNPCKLDQ256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpunpckldq256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
