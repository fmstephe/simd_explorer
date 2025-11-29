package vpbroadcastq

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastq_256.s
var assemblyVpbroadcastq256 string

//go:embed stub_vpbroadcastq_256.go
var stubVpbroadcastq256 string

type VPBROADCASTQ256 struct {
	scalar *number.Parameter
	ret    *number.Parameter
}

func NewVPBROADCASTQ256() *VPBROADCASTQ256 {
	return &VPBROADCASTQ256{
		scalar: number.NewNamedUintParameter("scalar", 64, 64, 16),
		ret:    number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPBROADCASTQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{v.scalar}
}

func (v *VPBROADCASTQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTQ256) Name() string {
	return "VPBROADCASTQ YMM (256 bit)"
}

func (v *VPBROADCASTQ256) Description() string {
	return "Broadcast 64-bit scalar to all lanes of YMM."
}

func (v *VPBROADCASTQ256) Stub() string {
	return stubVpbroadcastq256
}

func (v *VPBROADCASTQ256) Assembly() string {
	return assemblyVpbroadcastq256
}

func (v *VPBROADCASTQ256) Run(_ [][]byte) (output []byte) {
	q := number.ToUint64(v.scalar.FlatData())
	ret := [4]uint64{}
	vpbroadcastq256(q, &ret)
	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPBROADCASTQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
