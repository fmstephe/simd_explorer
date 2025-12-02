package vperm2i128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2i128_256_lowa_lowa.s
var assemblyVperm2i128256Lowa_lowa string

//go:embed stub_vperm2i128_256_lowa_lowa.go
var stubVperm2i128256Lowa_lowa string

type VPERM2I128256LOWA_LOWA struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2I128256LOWA_LOWA() *VPERM2I128256LOWA_LOWA {
	return &VPERM2I128256LOWA_LOWA{
		valsA: number.NewNamedUintParameter("valsA", 256, 64, 16),
		valsB: number.NewNamedUintParameter("valsB", 256, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPERM2I128256LOWA_LOWA) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2I128256LOWA_LOWA) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2I128256LOWA_LOWA) Name() string {
	return "VPERM2I128 (256 bit) lowa_lowa"
}

func (v *VPERM2I128256LOWA_LOWA) Description() string {
	return "dst.low = A.low, dst.high = A.low (per 128-bit lane)"
}

func (v *VPERM2I128256LOWA_LOWA) Stub() string {
	return stubVperm2i128256Lowa_lowa
}

func (v *VPERM2I128256LOWA_LOWA) Assembly() string {
	return assemblyVperm2i128256Lowa_lowa
}

func (v *VPERM2I128256LOWA_LOWA) Run() {
	valsA := [4]uint64{}
	copy(valsA[:], number.ToUint64Slice(v.valsA.FlatData()))
	valsB := [4]uint64{}
	copy(valsB[:], number.ToUint64Slice(v.valsB.FlatData()))

	ret := [4]uint64{}

	vperm2i128256Lowa_lowa(&valsA, &valsB, &ret)

	log.Printf("VPERM2I128256LOWA_LOWA valsA %v valsB %v ret %v", valsA, valsB, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERM2I128256LOWA_LOWA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
