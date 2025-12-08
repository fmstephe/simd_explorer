package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpckldq_128.s
var assemblyVpunpckldq128 string

//go:embed stub_vpunpckldq_128.go
var stubVpunpckldq128 string

type VPUNPCKLDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLDQ128() *VPUNPCKLDQ128 {
	return &VPUNPCKLDQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPUNPCKLDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLDQ128) Name() string {
	return "VPUNPCKLDQ (128 bit) "
}

func (v *VPUNPCKLDQ128) Description() string {
	return "Unpack and interleave low-order doublewords from two 128-bit sources."
}

func (v *VPUNPCKLDQ128) Stub() string {
	return stubVpunpckldq128
}

func (v *VPUNPCKLDQ128) Assembly() string {
	return assemblyVpunpckldq128
}

func (v *VPUNPCKLDQ128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpunpckldq128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLDQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
