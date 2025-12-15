package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_256_one.s
var assemblyVpslldq256One string

//go:embed stub_vpslldq_256_one.go
var stubVpslldq256One string

type VPSLLDQ256ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ256ONE() *VPSLLDQ256ONE {
	return &VPSLLDQ256ONE{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSLLDQ256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ256ONE) Name() string {
	return "VPSLLDQ (256 bit) one"
}

func (v *VPSLLDQ256ONE) Description() string {
	return "Shift left by 1 byte per 128-bit lane; lowest byte of each lane becomes zero."
}

func (v *VPSLLDQ256ONE) Stub() string {
	return stubVpslldq256One
}

func (v *VPSLLDQ256ONE) Assembly() string {
	return assemblyVpslldq256One
}

func (v *VPSLLDQ256ONE) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq256One(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
