package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_256_reverse.s
var assemblyVpermilpd256Reverse string

//go:embed stub_vpermilpd_256_reverse.go
var stubVpermilpd256Reverse string

type VPERMILPD256REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPD256REVERSE() *VPERMILPD256REVERSE {
	return &VPERMILPD256REVERSE{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VPERMILPD256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPD256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPD256REVERSE) Name() string {
	return "VPERMILPD (256 bit) reverse"
}

func (v *VPERMILPD256REVERSE) Description() string {
	return "Permute with imm8=0x1B per 128-bit lane: reverse lane order."
}

func (v *VPERMILPD256REVERSE) Stub() string {
	return stubVpermilpd256Reverse
}

func (v *VPERMILPD256REVERSE) Assembly() string {
	return assemblyVpermilpd256Reverse
}

func (v *VPERMILPD256REVERSE) Run() (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vpermilpd256Reverse(&vals, &ret)

	log.Printf("VPERMILPD256REVERSE vals %v ret %v", vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPD256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
