package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_512_k.s
var assemblyVpbroadcastb512K string

//go:embed stub_vpbroadcastb_512_k.go
var stubVpbroadcastb512K string

type VPBROADCASTB512K struct {
	scalar *number.Parameter
	pred   *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTB512K() *VPBROADCASTB512K {
	return &VPBROADCASTB512K{
		scalar: number.NewNamedUintParameter("scalar", 8, 8, 16),
		pred:   number.NewNamedUintParameter("pred", 64, 64, 16),
		ret:    number.NewNamedUintParameter("ret", 512, 64, 16),
	}
}

func (v *VPBROADCASTB512K) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar, v.pred}
}

func (v *VPBROADCASTB512K) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB512K) Name() string {
	return "VPBROADCASTB ZMM (512K bit) K"
}

func (v *VPBROADCASTB512K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to ZMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB512K) Stub() string {
	return stubVpbroadcastb512K
}

func (v *VPBROADCASTB512K) Assembly() string {
	return assemblyVpbroadcastb512K
}

func (v *VPBROADCASTB512K) Run() {
	// fields are initialized in constructor
	ret := [64]byte{}
	b := number.ToUint8(v.scalar.FlatData())
	k := number.ToUint64(v.pred.FlatData())
	vpbroadcastb512K(b, k, &ret)
	out := ret[:]
	log.Printf("VPBROADCASTB512K b %v k 0x%X ret %v", b, k, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTB512K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
