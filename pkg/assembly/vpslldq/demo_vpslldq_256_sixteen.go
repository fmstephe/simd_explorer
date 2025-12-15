package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_256_sixteen.s
var assemblyVpslldq256Sixteen string

//go:embed stub_vpslldq_256_sixteen.go
var stubVpslldq256Sixteen string

type VPSLLDQ256SIXTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ256SIXTEEN() *VPSLLDQ256SIXTEEN {
	return &VPSLLDQ256SIXTEEN{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSLLDQ256SIXTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ256SIXTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ256SIXTEEN) Name() string {
	return "VPSLLDQ (256 bit) sixteen"
}

func (v *VPSLLDQ256SIXTEEN) Description() string {
	return "Shift left by 16 bytes per 128-bit lane (full-lane zeroing)."
}

func (v *VPSLLDQ256SIXTEEN) Stub() string {
	return stubVpslldq256Sixteen
}

func (v *VPSLLDQ256SIXTEEN) Assembly() string {
	return assemblyVpslldq256Sixteen
}

func (v *VPSLLDQ256SIXTEEN) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq256Sixteen(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ256SIXTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
