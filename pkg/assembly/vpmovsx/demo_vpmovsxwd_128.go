package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxwd_128.s
var assemblyVpmovsxwd128 string

//go:embed stub_vpmovsxwd_128.go
var stubVpmovsxwd128 string

type VPMOVSXWD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXWD128() *VPMOVSXWD128 {
	return &VPMOVSXWD128{
		vals: number.NewNamedIntParameter("vals", 128, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMOVSXWD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXWD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXWD128) Name() string {
	return "VPMOVSXWD (128 bit) "
}

func (v *VPMOVSXWD128) Description() string {
	return "Sign-extend packed 16-bit integers to 32-bit integers."
}

func (v *VPMOVSXWD128) Stub() string {
	return stubVpmovsxwd128
}

func (v *VPMOVSXWD128) Assembly() string {
	return assemblyVpmovsxwd128
}

func (v *VPMOVSXWD128) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmovsxwd128(&vals, &ret)

	log.Printf("VPMOVSXWD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXWD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
