package punpckh

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckhdq_128.s
var assemblyVpunpckhdq128 string

//go:embed stub_vpunpckhdq_128.go
var stubVpunpckhdq128 string

type VPUNPCKHDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKHDQ128() *VPUNPCKHDQ128 {
	return &VPUNPCKHDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPUNPCKHDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKHDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKHDQ128) Name() string {
	return "VPUNPCKHDQ (128 bit) "
}

func (v *VPUNPCKHDQ128) Description() string {
	return "Unpack and interleave high-order doublewords from two 128-bit sources."
}

func (v *VPUNPCKHDQ128) Stub() string {
	return stubVpunpckhdq128
}

func (v *VPUNPCKHDQ128) Assembly() string {
	return assemblyVpunpckhdq128
}

func (v *VPUNPCKHDQ128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpunpckhdq128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKHDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKHDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
