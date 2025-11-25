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
}

func (v *VPBROADCASTD128) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(32, 32, 16)}
}

func (v *VPBROADCASTD128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
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

func (v *VPBROADCASTD128) Run(inputs [][]byte) (output []byte) {
	d := number.ToUint32(inputs[0])
	ret := [4]uint32{}
	vpbroadcastd128(d, &ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPBROADCASTD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
