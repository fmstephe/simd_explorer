package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhdq_256.s
var assemblyVpunpckhdq256 string

//go:embed stub_vpunpckhdq_256.go
var stubVpunpckhdq256 string

type VPUNPCKHDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHDQ256() *VPUNPCKHDQ256 {
	return &VPUNPCKHDQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPUNPCKHDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHDQ256) Name() string {
	return "VPUNPCKHDQ (256 bit) "
}

func (v *VPUNPCKHDQ256) Description() string {
	return "Unpack and interleave high-order doublewords from two 128-bit lanes of YMM inputs."
}

func (v *VPUNPCKHDQ256) Stub() string {
	return stubVpunpckhdq256
}

func (v *VPUNPCKHDQ256) Assembly() string {
	return assemblyVpunpckhdq256
}

func (v *VPUNPCKHDQ256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpunpckhdq256(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHDQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
