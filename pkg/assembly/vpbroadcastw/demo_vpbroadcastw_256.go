package vpbroadcastw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastw_256.s
var assemblyVpbroadcastw256 string

//go:embed stub_vpbroadcastw_256.go
var stubVpbroadcastw256 string

type VPBROADCASTW256 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTW256() *VPBROADCASTW256 {
	return &VPBROADCASTW256{
		scalar: number.NewNamedUintParameter("scalar", 16, 16, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPBROADCASTW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.scalar,
	}
}

func (v *VPBROADCASTW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTW256) Name() string {
	return "VPBROADCASTW YMM (256 bit)"
}

func (v *VPBROADCASTW256) Description() string {
	return "Broadcast 16-bit scalar to all lanes of YMM."
}

func (v *VPBROADCASTW256) Stub() string {
	return stubVpbroadcastw256
}

func (v *VPBROADCASTW256) Assembly() string {
	return assemblyVpbroadcastw256
}

func (v *VPBROADCASTW256) Run() {
	w := number.ToUint16(v.scalar.FlatData())
	ret := [16]uint16{}
	vpbroadcastw256(w, &ret)
	out := number.Uint16SliceToBytes(ret[:])
	log.Printf("VPBROADCASTW256 w 0x%X ret %v", w, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
