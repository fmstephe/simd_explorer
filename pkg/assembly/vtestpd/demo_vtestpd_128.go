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
}

func (v *VTESTPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),
		number.NewFloatParameter(128, 64),
	}
}

func (v *VTESTPD128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 16) // bit0=ZF, bit1=CF
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

func (v *VTESTPD128) Run(inputs [][]byte) (output []byte) {
	a := [2]float64{}
	copy(a[:], number.ToFloat64Slice(inputs[0]))
	b := [2]float64{}
	copy(b[:], number.ToFloat64Slice(inputs[1]))

	var flags uint32

	vtestpd128(&a, &b, &flags)

	log.Printf("VTESTPD128 A %v B %v flags 0x%X", a, b, flags)

	return number.Uint32ToBytes(flags)
}

func (v *VTESTPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
