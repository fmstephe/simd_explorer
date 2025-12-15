package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxdq_128.s
var assemblyVpmovsxdq128 string

//go:embed stub_vpmovsxdq_128.go
var stubVpmovsxdq128 string

type VPMOVSXDQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXDQ128() *VPMOVSXDQ128 {
	return &VPMOVSXDQ128{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVSXDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXDQ128) Name() string {
	return "VPMOVSXDQ (128 bit) "
}

func (v *VPMOVSXDQ128) Description() string {
	return "Sign-extend packed 32-bit integers to 64-bit integers."
}

func (v *VPMOVSXDQ128) Stub() string {
	return stubVpmovsxdq128
}

func (v *VPMOVSXDQ128) Assembly() string {
	return assemblyVpmovsxdq128
}

func (v *VPMOVSXDQ128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [2]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxdq128(&vals, &ret)

	log.Printf("VPMOVSXDQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
