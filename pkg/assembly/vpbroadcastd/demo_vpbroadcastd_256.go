package vpbroadcastd

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastd_256.s
var assemblyVpbroadcastd256 string

//go:embed stub_vpbroadcastd_256.go
var stubVpbroadcastd256 string

type VPBROADCASTD256 struct {
}

func (v *VPBROADCASTD256) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(32, 32, 16)}
}

func (v *VPBROADCASTD256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
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

func (v *VPBROADCASTD256) Run(inputs [][]byte) (output []byte) {
	d := number.ToUint32(inputs[0])
	ret := [8]uint32{}
	vpbroadcastd256(d, &ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VPBROADCASTD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
