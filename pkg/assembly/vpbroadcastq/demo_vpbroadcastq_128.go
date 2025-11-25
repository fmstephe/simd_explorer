package vpbroadcastq

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastq_128.s
var assemblyVpbroadcastq128 string

//go:embed stub_vpbroadcastq_128.go
var stubVpbroadcastq128 string

type VPBROADCASTQ128 struct {
}

func (v *VPBROADCASTQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(64, 64, 16)}
}

func (v *VPBROADCASTQ128) Output() *number.Parameter {
	return number.NewUintParameter(128, 64, 16)
}

func (v *VPBROADCASTQ128) Name() string {
	return "VPBROADCASTQ XMM (128 bit)"
}

func (v *VPBROADCASTQ128) Description() string {
	return "Broadcast 64-bit scalar to all lanes of XMM."
}

func (v *VPBROADCASTQ128) Stub() string {
	return stubVpbroadcastq128
}

func (v *VPBROADCASTQ128) Assembly() string {
	return assemblyVpbroadcastq128
}

func (v *VPBROADCASTQ128) Run(inputs [][]byte) (output []byte) {
	q := number.ToUint64(inputs[0])
	ret := [2]uint64{}
	vpbroadcastq128(q, &ret)
	return number.Uint64SliceToBytes(ret[:])
}

func (v *VPBROADCASTQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
