package vperm2i128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2i128_256_zeroed_zeroed.s
var assemblyVperm2i128256Zeroed_zeroed string

//go:embed stub_vperm2i128_256_zeroed_zeroed.go
var stubVperm2i128256Zeroed_zeroed string

type VPERM2I128256ZEROED_ZEROED struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2I128256ZEROED_ZEROED() *VPERM2I128256ZEROED_ZEROED {
	return &VPERM2I128256ZEROED_ZEROED{
		valsA: number.NewNamedUintParameter("valsA", 256, 64, 16),
		valsB: number.NewNamedUintParameter("valsB", 256, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPERM2I128256ZEROED_ZEROED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2I128256ZEROED_ZEROED) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2I128256ZEROED_ZEROED) Name() string {
	return "VPERM2I128 (256 bit) zeroed_zeroed"
}

func (v *VPERM2I128256ZEROED_ZEROED) Description() string {
	return "dst.low = zero, dst.high = zero (per 128-bit lane)"
}

func (v *VPERM2I128256ZEROED_ZEROED) Stub() string {
	return stubVperm2i128256Zeroed_zeroed
}

func (v *VPERM2I128256ZEROED_ZEROED) Assembly() string {
	return assemblyVperm2i128256Zeroed_zeroed
}

func (v *VPERM2I128256ZEROED_ZEROED) Run() {
	valsA := [4]uint64{}
	copy(valsA[:], number.ToUint64Slice(v.valsA.FlatData()))
	valsB := [4]uint64{}
	copy(valsB[:], number.ToUint64Slice(v.valsB.FlatData()))

	ret := [4]uint64{}

	vperm2i128256Zeroed_zeroed(&valsA, &valsB, &ret)

	log.Printf("VPERM2I128256ZEROED_ZEROED valsA %v valsB %v ret %v", valsA, valsB, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERM2I128256ZEROED_ZEROED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
