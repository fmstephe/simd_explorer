package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbq_256.s
var assemblyVpmovsxbq256 string

//go:embed stub_vpmovsxbq_256.go
var stubVpmovsxbq256 string

type VPMOVSXBQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBQ256() *VPMOVSXBQ256 {
	return &VPMOVSXBQ256{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVSXBQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBQ256) Name() string {
	return "VPMOVSXBQ (256 bit) "
}

func (v *VPMOVSXBQ256) Description() string {
	return "Sign-extend packed 8-bit integers to 64-bit integers, per 128-bit lane."
}

func (v *VPMOVSXBQ256) Stub() string {
	return stubVpmovsxbq256
}

func (v *VPMOVSXBQ256) Assembly() string {
	return assemblyVpmovsxbq256
}

func (v *VPMOVSXBQ256) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [4]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxbq256(&vals, &ret)

	log.Printf("VPMOVSXBQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
