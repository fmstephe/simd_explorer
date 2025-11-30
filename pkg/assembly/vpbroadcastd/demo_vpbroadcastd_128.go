package vpbroadcastd

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastd_128.s
var assemblyVpbroadcastd128 string

//go:embed stub_vpbroadcastd_128.go
var stubVpbroadcastd128 string

type VPBROADCASTD128 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTD128() *VPBROADCASTD128 {
	return &VPBROADCASTD128{
		scalar: number.NewNamedUintParameter("scalar", 32, 32, 16),
		ret:    number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPBROADCASTD128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar}
}

func (v *VPBROADCASTD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTD128) Name() string {
	return "VPBROADCASTD XMM (128 bit)"
}

func (v *VPBROADCASTD128) Description() string {
	return "Broadcast 32-bit scalar to all lanes of XMM."
}

func (v *VPBROADCASTD128) Stub() string {
	return stubVpbroadcastd128
}

func (v *VPBROADCASTD128) Assembly() string {
	return assemblyVpbroadcastd128
}

func (v *VPBROADCASTD128) Run() (output []byte) {
	d := number.ToUint32(v.scalar.FlatData())
	ret := [4]uint32{}
	vpbroadcastd128(d, &ret)
	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPBROADCASTD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
