package vpbroadcastw

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastw_128.s
var assemblyVpbroadcastw128 string

//go:embed stub_vpbroadcastw_128.go
var stubVpbroadcastw128 string

type VPBROADCASTW128 struct {
}

func (v *VPBROADCASTW128) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(16, 16, 16)}
}

func (v *VPBROADCASTW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 16)
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

func (v *VPBROADCASTW128) Run(inputs [][]byte) (output []byte) {
	w := number.ToUint16(inputs[0])
	ret := [8]uint16{}
	vpbroadcastw128(w, &ret)
	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPBROADCASTW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
