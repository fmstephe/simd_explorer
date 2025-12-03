package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_512.s
var assemblyVpbroadcastb512 string

//go:embed stub_vpbroadcastb_512.go
var stubVpbroadcastb512 string

type VPBROADCASTB512 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTB512() *VPBROADCASTB512 {
	return &VPBROADCASTB512{
		scalar: number.NewNamedUintParameter("scalar", 8, 8, 10),
		ret:    number.NewNamedUintParameter("ret", 512, 64, 10),
	}
}

func (v *VPBROADCASTB512) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
	}
}

func (v *VPBROADCASTB512) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB512) Name() string {
	return "VPBROADCASTB ZMM (512 bit)"
}

func (v *VPBROADCASTB512) Description() string {
	return "Broadcast an 8-bit value to all 64 byte elements in ZMM."
}

func (v *VPBROADCASTB512) Stub() string {
	return stubVpbroadcastb512
}

func (v *VPBROADCASTB512) Assembly() string {
	return assemblyVpbroadcastb512
}

func (v *VPBROADCASTB512) Run() {
	ret := [64]byte{}
	b := number.ToUint8(v.scalar.FlatData())
	vpbroadcastb512(b, &ret)
	out := ret[:]
	log.Printf("VPBROADCASTB512 b %v ret %v", b, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTB512) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
