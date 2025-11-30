package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_256_k.s
var assemblyVpbroadcastb256K string

//go:embed stub_vpbroadcastb_256_k.go
var stubVpbroadcastb256K string

type VPBROADCASTB256K struct {
	scalar *number.Parameter
	pred   *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTB256K() *VPBROADCASTB256K {
	return &VPBROADCASTB256K{
		scalar: number.NewNamedUintParameter("scalar", 8, 8, 16),
		pred:   number.NewNamedUintParameter("pred", 64, 64, 16),
		ret:    number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPBROADCASTB256K) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar, v.pred}
}

func (v *VPBROADCASTB256K) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB256K) Name() string {
	return "VPBROADCASTB YMM (256K bit) K"
}

func (v *VPBROADCASTB256K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to YMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB256K) Stub() string {
	return stubVpbroadcastb256K
}

func (v *VPBROADCASTB256K) Assembly() string {
	return assemblyVpbroadcastb256K
}

func (v *VPBROADCASTB256K) Run() {
	ret := [32]byte{}
	b := number.ToUint8(v.scalar.FlatData())
	k := number.ToUint64(v.pred.FlatData())
	vpbroadcastb256K(b, k, &ret)
	out := ret[:]
	log.Printf("VPBROADCASTB256K b %v k 0x%X ret %v", b, k, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTB256K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
