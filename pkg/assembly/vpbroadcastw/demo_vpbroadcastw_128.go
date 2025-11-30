package vpbroadcastw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastw_128.s
var assemblyVpbroadcastw128 string

//go:embed stub_vpbroadcastw_128.go
var stubVpbroadcastw128 string

type VPBROADCASTW128 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTW128() *VPBROADCASTW128 {
	return &VPBROADCASTW128{
		scalar: number.NewNamedUintParameter("scalar", 16, 16, 16),
		ret:    number.NewNamedUintParameter("ret", 128, 16, 16),
	}
}

func (v *VPBROADCASTW128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar}
}

func (v *VPBROADCASTW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTW128) Name() string {
	return "VPBROADCASTW XMM (128 bit)"
}

func (v *VPBROADCASTW128) Description() string {
	return "Broadcast 16-bit scalar to all lanes of XMM."
}

func (v *VPBROADCASTW128) Stub() string {
	return stubVpbroadcastw128
}

func (v *VPBROADCASTW128) Assembly() string {
	return assemblyVpbroadcastw128
}

func (v *VPBROADCASTW128) Run() {
	w := number.ToUint16(v.scalar.FlatData())
	ret := [8]uint16{}
	vpbroadcastw128(w, &ret)
	out := number.Uint16SliceToBytes(ret[:])
	log.Printf("VPBROADCASTW128 w 0x%X ret %v", w, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
