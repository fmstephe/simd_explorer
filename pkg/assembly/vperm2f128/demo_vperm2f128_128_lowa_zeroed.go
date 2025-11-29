package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowa_zeroed.s
var assemblyVperm2f128128Lowa_zeroed string

//go:embed stub_vperm2f128_128_lowa_zeroed.go
var stubVperm2f128128Lowa_zeroed string

type VPERM2F128128LOWA_ZEROED struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2F128128LOWA_ZEROED() *VPERM2F128128LOWA_ZEROED {
	return &VPERM2F128128LOWA_ZEROED{
		valsA: number.NewNamedFloatParameter("valsA", 256, 32),
		valsB: number.NewNamedFloatParameter("valsB", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERM2F128128LOWA_ZEROED) Inputs() []*number.Parameter {
	return []*number.Parameter{v.valsA, v.valsB}
}

func (v *VPERM2F128128LOWA_ZEROED) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2F128128LOWA_ZEROED) Name() string {
	return "VPERM2F128 (128 bit) lowa_zeroed"
}

func (v *VPERM2F128128LOWA_ZEROED) Description() string {
	return "dst.low = A.low, dst.high = zero (per 128-bit lane)"
}

func (v *VPERM2F128128LOWA_ZEROED) Stub() string {
	return stubVperm2f128128Lowa_zeroed
}

func (v *VPERM2F128128LOWA_ZEROED) Assembly() string {
	return assemblyVperm2f128128Lowa_zeroed
}

func (v *VPERM2F128128LOWA_ZEROED) Run(_ [][]byte) (output []byte) {
	valsA := [8]float32{}
	copy(valsA[:], number.ToFloat32Slice(v.valsA.FlatData()))
	valsB := [8]float32{}
	copy(valsB[:], number.ToFloat32Slice(v.valsB.FlatData()))
	ret := [8]float32{}

	vperm2f128128Lowa_zeroed(&valsA, &valsB, &ret)

	log.Printf("VPERM2F128128LOWA_ZEROED A %v B %v output %v", valsA, valsB, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERM2F128128LOWA_ZEROED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
