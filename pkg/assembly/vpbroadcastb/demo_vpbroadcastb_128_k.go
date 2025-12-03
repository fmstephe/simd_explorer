package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_128_k.s
var assemblyVpbroadcastb128K string

//go:embed stub_vpbroadcastb_128_k.go
var stubVpbroadcastb128K string

type VPBROADCASTB128K struct {
	scalar *number.Parameter
	pred   *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTB128K() *VPBROADCASTB128K {
	return &VPBROADCASTB128K{
		scalar: number.NewNamedUintParameter("scalar", 8, 8, 10),
		pred:   number.NewNamedUintParameter("pred", 64, 64, 16),
		ret:    number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPBROADCASTB128K) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
		v.pred,
	}
}

func (v *VPBROADCASTB128K) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB128K) Name() string {
	return "VPBROADCASTB XMM (128K bit) K"
}

func (v *VPBROADCASTB128K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to XMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB128K) Stub() string {
	return stubVpbroadcastb128K
}

func (v *VPBROADCASTB128K) Assembly() string {
	return assemblyVpbroadcastb128K
}

func (v *VPBROADCASTB128K) Run() {
	ret := [16]byte{}
	b := number.ToUint8(v.scalar.FlatData())
	k := number.ToUint64(v.pred.FlatData())
	vpbroadcastb128K(b, k, &ret)
	out := ret[:]
	log.Printf("VPBROADCASTB128K b %v k 0x%X ret %v", b, k, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTB128K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
