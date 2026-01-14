package vpbroadcastb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_256.s
var assemblyVpbroadcastb256 string

//go:embed stub_vpbroadcastb_256.go
var stubVpbroadcastb256 string

type VPBROADCASTB256 struct {
	b   *number.Parameter
	ret *number.Parameter
}

func NewVPBROADCASTB256() *VPBROADCASTB256 {
	return &VPBROADCASTB256{
		b:   number.NewNamedUintParameter("b", 8, 8, 10),
		ret: number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPBROADCASTB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.b,
	}
}

func (v *VPBROADCASTB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPBROADCASTB256) Name() string {
	return "VPBROADCASTB (256 bit)"
}

func (v *VPBROADCASTB256) Description() string {
	return "Broadcast an 8-bit value to all 32 byte elements in YMM."
}

func (v *VPBROADCASTB256) Stub() string {
	return stubVpbroadcastb256
}

func (v *VPBROADCASTB256) Assembly() string {
	return assemblyVpbroadcastb256
}

func (v *VPBROADCASTB256) Run() {
	b := number.ToUint8(v.b.FlatData())
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpbroadcastb256(b, &ret)

	log.Printf("VPBROADCASTB256 b %v ret %v", b, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPBROADCASTB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
