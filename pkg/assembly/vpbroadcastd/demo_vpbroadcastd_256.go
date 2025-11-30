package vpbroadcastd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastd_256.s
var assemblyVpbroadcastd256 string

//go:embed stub_vpbroadcastd_256.go
var stubVpbroadcastd256 string

type VPBROADCASTD256 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTD256() *VPBROADCASTD256 {
	return &VPBROADCASTD256{
		scalar: number.NewNamedUintParameter("scalar", 32, 32, 16),
		ret:    number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VPBROADCASTD256) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar}
}

func (v *VPBROADCASTD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTD256) Name() string {
	return "VPBROADCASTD YMM (256 bit)"
}

func (v *VPBROADCASTD256) Description() string {
	return "Broadcast 32-bit scalar to all lanes of YMM."
}

func (v *VPBROADCASTD256) Stub() string {
	return stubVpbroadcastd256
}

func (v *VPBROADCASTD256) Assembly() string {
	return assemblyVpbroadcastd256
}

func (v *VPBROADCASTD256) Run() {
	d := number.ToUint32(v.scalar.FlatData())
	ret := [8]uint32{}
	vpbroadcastd256(d, &ret)
	out := number.Uint32SliceToBytes(ret[:])
	log.Printf("VPBROADCASTD256 d 0x%X ret %v", d, ret)
	v.ret.SetData(out)

}

func (v *VPBROADCASTD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
