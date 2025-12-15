package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxwq_256.s
var assemblyVpmovsxwq256 string

//go:embed stub_vpmovsxwq_256.go
var stubVpmovsxwq256 string

type VPMOVSXWQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXWQ256() *VPMOVSXWQ256 {
	return &VPMOVSXWQ256{
		vals: number.NewNamedIntParameter("vals", 128, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVSXWQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXWQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXWQ256) Name() string {
	return "VPMOVSXWQ (256 bit) "
}

func (v *VPMOVSXWQ256) Description() string {
	return "Sign-extend packed 16-bit integers to 64-bit integers, per 128-bit lane."
}

func (v *VPMOVSXWQ256) Stub() string {
	return stubVpmovsxwq256
}

func (v *VPMOVSXWQ256) Assembly() string {
	return assemblyVpmovsxwq256
}

func (v *VPMOVSXWQ256) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [4]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxwq256(&vals, &ret)

	log.Printf("VPMOVSXWQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXWQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
