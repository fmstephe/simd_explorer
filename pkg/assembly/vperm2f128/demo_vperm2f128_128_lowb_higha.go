package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowb_higha.s
var assemblyVperm2f128128Lowb_higha string

//go:embed stub_vperm2f128_128_lowb_higha.go
var stubVperm2f128128Lowb_higha string

type VPERM2F128128LOWB_HIGHA struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128128LOWB_HIGHA() *VPERM2F128128LOWB_HIGHA {
	return &VPERM2F128128LOWB_HIGHA{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128128LOWB_HIGHA) Inputs() []*number.Parameter {
	return []*number.Parameter{v.valsA, v.valsB}
}

func (v *VPERM2F128128LOWB_HIGHA) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128128LOWB_HIGHA) Name() string {
	return "VPERM2F128 (128 bit) lowb_higha"
}

func (v *VPERM2F128128LOWB_HIGHA) Description() string {
	return "dst.low = B.low, dst.high = A.high (per 128-bit lane)"
}

func (v *VPERM2F128128LOWB_HIGHA) Stub() string {
	return stubVperm2f128128Lowb_higha
}

func (v *VPERM2F128128LOWB_HIGHA) Assembly() string {
	return assemblyVperm2f128128Lowb_higha
}

func (v *VPERM2F128128LOWB_HIGHA) Run(_ [][]byte) (output []byte) {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128128Lowb_higha(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128128LOWB_HIGHA A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERM2F128128LOWB_HIGHA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
