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
}

func (v *VTESTPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),
		number.NewFloatParameter(256, 64),
	}
}

func (v *VTESTPD256) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 16) // bit0=ZF, bit1=CF
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

func (v *VTESTPD256) Run(inputs [][]byte) (output []byte) {
	a := [4]float64{}
	copy(a[:], number.ToFloat64Slice(inputs[0]))
	b := [4]float64{}
	copy(b[:], number.ToFloat64Slice(inputs[1]))

	var flags uint32

	vtestpd256(&a, &b, &flags)

	log.Printf("VTESTPD256 A %v B %v flags 0x%X", a, b, flags)

	return number.Uint32ToBytes(flags)
}

func (v *VTESTPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
