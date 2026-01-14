package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_128.s
var assemblyVpbroadcastb128 string

//go:embed stub_vpbroadcastb_128.go
var stubVpbroadcastb128 string

type VPBROADCASTB128 struct {
	b   *number.Parameter
	ret *number.Parameter
}

func NewVPBROADCASTB128() *VPBROADCASTB128 {
	return &VPBROADCASTB128{
		b:   number.NewNamedUintParameter("b", 8, 8, 10),
		ret: number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPBROADCASTB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.b,
	}
}

func (v *VPBROADCASTB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB128) Name() string {
	return "VPBROADCASTB (128 bit)"
}

func (v *VPBROADCASTB128) Description() string {
	return "Broadcast an 8-bit value to all 16 byte elements in XMM."
}

func (v *VPBROADCASTB128) Stub() string {
	return stubVpbroadcastb128
}

func (v *VPBROADCASTB128) Assembly() string {
	return assemblyVpbroadcastb128
}

func (v *VPBROADCASTB128) Run() {
	b := number.ToUint8(v.b.FlatData())
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpbroadcastb128(b, &ret)

	log.Printf("VPBROADCASTB128 b %v ret %v", b, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPBROADCASTB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
