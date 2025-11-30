package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowa_lowa.s
var assemblyVperm2f128128Lowa_lowa string

//go:embed stub_vperm2f128_128_lowa_lowa.go
var stubVperm2f128128Lowa_lowa string

type VPERM2F128128LOWA_LOWA struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128128LOWA_LOWA() *VPERM2F128128LOWA_LOWA {
	return &VPERM2F128128LOWA_LOWA{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128128LOWA_LOWA) Inputs() []*number.Parameter {
	return []*number.Parameter{v.valsA, v.valsB}
}

func (v *VPERM2F128128LOWA_LOWA) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128128LOWA_LOWA) Name() string {
	return "VPERM2F128 (128 bit) lowa_lowa"
}

func (v *VPERM2F128128LOWA_LOWA) Description() string {
	return "dst.low = A.low, dst.high = A.low (per 128-bit lane)"
}

func (v *VPERM2F128128LOWA_LOWA) Stub() string {
	return stubVperm2f128128Lowa_lowa
}

func (v *VPERM2F128128LOWA_LOWA) Assembly() string {
	return assemblyVperm2f128128Lowa_lowa
}

func (v *VPERM2F128128LOWA_LOWA) Run() {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128128Lowa_lowa(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128128LOWA_LOWA A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERM2F128128LOWA_LOWA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
