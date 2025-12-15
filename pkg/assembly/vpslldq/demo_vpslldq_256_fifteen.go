package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_256_fifteen.s
var assemblyVpslldq256Fifteen string

//go:embed stub_vpslldq_256_fifteen.go
var stubVpslldq256Fifteen string

type VPSLLDQ256FIFTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ256FIFTEEN() *VPSLLDQ256FIFTEEN {
	return &VPSLLDQ256FIFTEEN{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSLLDQ256FIFTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ256FIFTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ256FIFTEEN) Name() string {
	return "VPSLLDQ (256 bit) fifteen"
}

func (v *VPSLLDQ256FIFTEEN) Description() string {
	return "Shift left by 15 bytes per 128-bit lane; only top byte of each lane may survive."
}

func (v *VPSLLDQ256FIFTEEN) Stub() string {
	return stubVpslldq256Fifteen
}

func (v *VPSLLDQ256FIFTEEN) Assembly() string {
	return assemblyVpslldq256Fifteen
}

func (v *VPSLLDQ256FIFTEEN) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq256Fifteen(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ256FIFTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
