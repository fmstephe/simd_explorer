package vtestpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vtestpd_256.s
var assemblyVtestpd256 string

//go:embed stub_vtestpd_256.go
var stubVtestpd256 string

type VTESTPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVTESTPD256() *VTESTPD256 {
	return &VTESTPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VTESTPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VTESTPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VTESTPD256) Name() string {
	return "VTESTPD (256 bit) "
}

func (v *VTESTPD256) Description() string {
	return "Test packed double-precision sign bits (per 128-bit lane). ZF=1 if (a & b) sign bits all zero; CF=1 if (~a & b) sign bits all zero. Output: bit0=ZF, bit1=CF."
}

func (v *VTESTPD256) Stub() string {
	return stubVtestpd256
}

func (v *VTESTPD256) Assembly() string {
	return assemblyVtestpd256
}

func (v *VTESTPD256) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))

	var flags uint32

	vtestpd256(&vals1, &vals2, &flags)

	log.Printf("VTESTPD256 vals1 %v vals2 %v flags 0x%X", vals1, vals2, flags)

	out := number.Uint32ToBytes(flags)
	v.ret.SetData(out)
	return out
}

func (v *VTESTPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
