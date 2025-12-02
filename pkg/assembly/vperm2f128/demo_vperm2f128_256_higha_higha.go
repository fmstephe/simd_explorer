package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_256_higha_higha.s
var assemblyVperm2f128256Higha_higha string

//go:embed stub_vperm2f128_256_higha_higha.go
var stubVperm2f128256Higha_higha string

type VPERM2F128256HIGHA_HIGHA struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128256HIGHA_HIGHA() *VPERM2F128256HIGHA_HIGHA {
	return &VPERM2F128256HIGHA_HIGHA{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128256HIGHA_HIGHA) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2F128256HIGHA_HIGHA) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128256HIGHA_HIGHA) Name() string {
	return "VPERM2F128 (128 bit) higha_higha"
}

func (v *VPERM2F128256HIGHA_HIGHA) Description() string {
	return "dst.low = A.high, dst.high = A.high (per 128-bit lane)"
}

func (v *VPERM2F128256HIGHA_HIGHA) Stub() string {
	return stubVperm2f128256Higha_higha
}

func (v *VPERM2F128256HIGHA_HIGHA) Assembly() string {
	return assemblyVperm2f128256Higha_higha
}

func (v *VPERM2F128256HIGHA_HIGHA) Run() {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128256Higha_higha(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128256HIGHA_HIGHA A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERM2F128256HIGHA_HIGHA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
