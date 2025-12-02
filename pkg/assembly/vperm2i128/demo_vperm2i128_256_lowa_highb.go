package vperm2i128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2i128_256_lowa_highb.s
var assemblyVperm2i128256Lowa_highb string

//go:embed stub_vperm2i128_256_lowa_highb.go
var stubVperm2i128256Lowa_highb string

type VPERM2I128256LOWA_HIGHB struct {
	valsA *number.Parameter
	valsB *number.Parameter
	ret   *number.Parameter
}

func NewVPERM2I128256LOWA_HIGHB() *VPERM2I128256LOWA_HIGHB {
	return &VPERM2I128256LOWA_HIGHB{
		valsA: number.NewNamedUintParameter("valsA", 256, 64, 16),
		valsB: number.NewNamedUintParameter("valsB", 256, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPERM2I128256LOWA_HIGHB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.valsA,
		v.valsB,
	}
}

func (v *VPERM2I128256LOWA_HIGHB) Output() *number.Parameter {
	return v.ret
}

func (v *VPERM2I128256LOWA_HIGHB) Name() string {
	return "VPERM2I128 (256 bit) lowa_highb"
}

func (v *VPERM2I128256LOWA_HIGHB) Description() string {
	return "dst.low = A.low, dst.high = B.high (per 128-bit lane)"
}

func (v *VPERM2I128256LOWA_HIGHB) Stub() string {
	return stubVperm2i128256Lowa_highb
}

func (v *VPERM2I128256LOWA_HIGHB) Assembly() string {
	return assemblyVperm2i128256Lowa_highb
}

func (v *VPERM2I128256LOWA_HIGHB) Run() {
	valsA := [4]uint64{}
	copy(valsA[:], number.ToUint64Slice(v.valsA.FlatData()))
	valsB := [4]uint64{}
	copy(valsB[:], number.ToUint64Slice(v.valsB.FlatData()))

	ret := [4]uint64{}

	vperm2i128256Lowa_highb(&valsA, &valsB, &ret)

	log.Printf("VPERM2I128256LOWA_HIGHB valsA %v valsB %v ret %v", valsA, valsB, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERM2I128256LOWA_HIGHB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
