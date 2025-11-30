package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowb_highb.s
var assemblyVperm2f128128Lowb_highb string

//go:embed stub_vperm2f128_128_lowb_highb.go
var stubVperm2f128128Lowb_highb string

type VPERM2F128128LOWB_HIGHB struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128128LOWB_HIGHB() *VPERM2F128128LOWB_HIGHB {
	return &VPERM2F128128LOWB_HIGHB{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128128LOWB_HIGHB) Inputs() []*number.Parameter {
	return []*number.Parameter{v.valsA, v.valsB}
}

func (v *VPERM2F128128LOWB_HIGHB) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128128LOWB_HIGHB) Name() string {
	return "VPERM2F128 (128 bit) lowb_highb"
}

func (v *VPERM2F128128LOWB_HIGHB) Description() string {
	return "dst.low = B.low, dst.high = B.high (per 128-bit lane)"
}

func (v *VPERM2F128128LOWB_HIGHB) Stub() string {
	return stubVperm2f128128Lowb_highb
}

func (v *VPERM2F128128LOWB_HIGHB) Assembly() string {
	return assemblyVperm2f128128Lowb_highb
}

func (v *VPERM2F128128LOWB_HIGHB) Run() {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128128Lowb_highb(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128128LOWB_HIGHB A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERM2F128128LOWB_HIGHB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
