package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbw_256.s
var assemblyVpmovsxbw256 string

//go:embed stub_vpmovsxbw_256.go
var stubVpmovsxbw256 string

type VPMOVSXBW256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBW256() *VPMOVSXBW256 {
	return &VPMOVSXBW256{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMOVSXBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBW256) Name() string {
	return "VPMOVSXBW (256 bit) "
}

func (v *VPMOVSXBW256) Description() string {
	return "Sign-extend packed 8-bit integers to 16-bit integers, per 128-bit lane."
}

func (v *VPMOVSXBW256) Stub() string {
	return stubVpmovsxbw256
}

func (v *VPMOVSXBW256) Assembly() string {
	return assemblyVpmovsxbw256
}

func (v *VPMOVSXBW256) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [16]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpmovsxbw256(&vals, &ret)

	log.Printf("VPMOVSXBW vals %v ret %v", vals, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
