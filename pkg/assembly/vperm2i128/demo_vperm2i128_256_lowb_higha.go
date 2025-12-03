package vperm2i128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2i128_256_lowb_higha.s
var assemblyVperm2i128256Lowb_higha string

//go:embed stub_vperm2i128_256_lowb_higha.go
var stubVperm2i128256Lowb_higha string

type VPERM2I128256LOWB_HIGHA struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2I128256LOWB_HIGHA() *VPERM2I128256LOWB_HIGHA {
	return &VPERM2I128256LOWB_HIGHA{
		valsA: number.NewNamedUintParameter("valsA", 256, 64, 10),
		valsB: number.NewNamedUintParameter("valsB", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERM2I128256LOWB_HIGHA) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2I128256LOWB_HIGHA) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2I128256LOWB_HIGHA) Name() string {
	return "VPERM2I128 (256 bit) lowb_higha"
}

func (v *VPERM2I128256LOWB_HIGHA) Description() string {
	return "dst.low = B.low, dst.high = A.high (per 128-bit lane)"
}

func (v *VPERM2I128256LOWB_HIGHA) Stub() string {
	return stubVperm2i128256Lowb_higha
}

func (v *VPERM2I128256LOWB_HIGHA) Assembly() string {
	return assemblyVperm2i128256Lowb_higha
}

func (v *VPERM2I128256LOWB_HIGHA) Run() {
	valsA := [4]uint64{}
	copy(valsA[:], number.ToUint64Slice(v.valsA.FlatData()))
	valsB := [4]uint64{}
	copy(valsB[:], number.ToUint64Slice(v.valsB.FlatData()))

	ret := [4]uint64{}

	vperm2i128256Lowb_higha(&valsA, &valsB, &ret)

	log.Printf("VPERM2I128256LOWB_HIGHA valsA %v valsB %v ret %v", valsA, valsB, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERM2I128256LOWB_HIGHA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
