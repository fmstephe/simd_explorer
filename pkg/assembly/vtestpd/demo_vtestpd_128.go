package vtestpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vtestpd_128.s
var assemblyVtestpd128 string

//go:embed stub_vtestpd_128.go
var stubVtestpd128 string

type VTESTPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVTESTPD128() *VTESTPD128 {
	return &VTESTPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VTESTPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VTESTPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VTESTPD128) Name() string {
	return "VTESTPD (128 bit) "
}

func (v *VTESTPD128) Description() string {
	return "Test packed double-precision sign bits. ZF=1 if (a & b) sign bits all zero; CF=1 if (~a & b) sign bits all zero. Output: bit0=ZF, bit1=CF."
}

func (v *VTESTPD128) Stub() string {
	return stubVtestpd128
}

func (v *VTESTPD128) Assembly() string {
	return assemblyVtestpd128
}

func (v *VTESTPD128) Run() (output []byte) {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))

	var flags uint32

	vtestpd128(&vals1, &vals2, &flags)

	log.Printf("VTESTPD128 vals1 %v vals2 %v flags 0x%X", vals1, vals2, flags)

	out := number.Uint32ToBytes(flags)
	v.ret.SetData(out)
	return out
}

func (v *VTESTPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
