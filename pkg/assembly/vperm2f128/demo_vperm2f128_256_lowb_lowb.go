package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_256_lowb_lowb.s
var assemblyVperm2f128256Lowb_lowb string

//go:embed stub_vperm2f128_256_lowb_lowb.go
var stubVperm2f128256Lowb_lowb string

type VPERM2F128256LOWB_LOWB struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128256LOWB_LOWB() *VPERM2F128256LOWB_LOWB {
	return &VPERM2F128256LOWB_LOWB{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128256LOWB_LOWB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2F128256LOWB_LOWB) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128256LOWB_LOWB) Name() string {
	return "VPERM2F128 (128 bit) lowb_lowb"
}

func (v *VPERM2F128256LOWB_LOWB) Description() string {
	return "dst.low = B.low, dst.high = B.low (per 128-bit lane)"
}

func (v *VPERM2F128256LOWB_LOWB) Stub() string {
	return stubVperm2f128256Lowb_lowb
}

func (v *VPERM2F128256LOWB_LOWB) Assembly() string {
	return assemblyVperm2f128256Lowb_lowb
}

func (v *VPERM2F128256LOWB_LOWB) Run() {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128256Lowb_lowb(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128256LOWB_LOWB A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERM2F128256LOWB_LOWB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
