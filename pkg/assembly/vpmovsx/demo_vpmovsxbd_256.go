package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbd_256.s
var assemblyVpmovsxbd256 string

//go:embed stub_vpmovsxbd_256.go
var stubVpmovsxbd256 string

type VPMOVSXBD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBD256() *VPMOVSXBD256 {
	return &VPMOVSXBD256{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMOVSXBD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBD256) Name() string {
	return "VPMOVSXBD (256 bit) "
}

func (v *VPMOVSXBD256) Description() string {
	return "Sign-extend packed 8-bit integers to 32-bit integers, per 128-bit lane."
}

func (v *VPMOVSXBD256) Stub() string {
	return stubVpmovsxbd256
}

func (v *VPMOVSXBD256) Assembly() string {
	return assemblyVpmovsxbd256
}

func (v *VPMOVSXBD256) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmovsxbd256(&vals, &ret)

	log.Printf("VPMOVSXBD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
