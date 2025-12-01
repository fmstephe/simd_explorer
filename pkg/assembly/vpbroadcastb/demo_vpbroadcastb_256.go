package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_256.s
var assemblyVpbroadcastb256 string

//go:embed stub_vpbroadcastb_256.go
var stubVpbroadcastb256 string

type VPBROADCASTB256 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

// constructor defined at bottom

func (v *VPBROADCASTB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
	}
}

func (v *VPBROADCASTB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB256) Name() string {
	return "VPBROADCASTB YMM (256 bit)"
}

func (v *VPBROADCASTB256) Description() string {
	return "Broadcast an 8-bit value to all 32 byte elements in YMM."
}

func (v *VPBROADCASTB256) Stub() string {
	return stubVpbroadcastb256
}

func (v *VPBROADCASTB256) Assembly() string {
	return assemblyVpbroadcastb256
}

func (v *VPBROADCASTB256) Run() {
	ret := [32]byte{}
	b := number.ToUint8(v.scalar.FlatData())
	vpbroadcastb256(b, &ret)
	out := ret[:]
	log.Printf("VPBROADCASTB256 b %v ret %v", b, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}

func NewVPBROADCASTB256() *VPBROADCASTB256 {
	v := &VPBROADCASTB256{
		scalar: number.NewNamedUintParameter("scalar", 8, 8, 16),
		ret:    number.NewNamedUintParameter("ret", 256, 64, 16),
	}
	return v
}
