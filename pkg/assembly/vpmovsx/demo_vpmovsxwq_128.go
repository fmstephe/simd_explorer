package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxwq_128.s
var assemblyVpmovsxwq128 string

//go:embed stub_vpmovsxwq_128.go
var stubVpmovsxwq128 string

type VPMOVSXWQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXWQ128() *VPMOVSXWQ128 {
	return &VPMOVSXWQ128{
		vals: number.NewNamedIntParameter("vals", 128, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVSXWQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXWQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXWQ128) Name() string {
	return "VPMOVSXWQ (128 bit) "
}

func (v *VPMOVSXWQ128) Description() string {
	return "Sign-extend packed 16-bit integers to 64-bit integers."
}

func (v *VPMOVSXWQ128) Stub() string {
	return stubVpmovsxwq128
}

func (v *VPMOVSXWQ128) Assembly() string {
	return assemblyVpmovsxwq128
}

func (v *VPMOVSXWQ128) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [2]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxwq128(&vals, &ret)

	log.Printf("VPMOVSXWQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXWQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
