package vpbroadcastw

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastw_256.s
var assemblyVpbroadcastw256 string

//go:embed stub_vpbroadcastw_256.go
var stubVpbroadcastw256 string

type VPBROADCASTW256 struct {
}

func (v *VPBROADCASTW256) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(16, 16, 16)}
}

func (v *VPBROADCASTW256) Output() *number.Parameter {
	return number.NewUintParameter(256, 16, 16)
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

func (v *VPBROADCASTW256) Run(inputs [][]byte) (output []byte) {
	w := number.ToUint16(inputs[0])
	ret := [16]uint16{}
	vpbroadcastw256(w, &ret)
	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPBROADCASTW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
