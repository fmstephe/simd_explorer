package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxdq_256.s
var assemblyVpmovsxdq256 string

//go:embed stub_vpmovsxdq_256.go
var stubVpmovsxdq256 string

type VPMOVSXDQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXDQ256() *VPMOVSXDQ256 {
	return &VPMOVSXDQ256{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVSXDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXDQ256) Name() string {
	return "VPMOVSXDQ (256 bit) "
}

func (v *VPMOVSXDQ256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVSXDQ256) Stub() string {
	return stubVpmovsxdq256
}

func (v *VPMOVSXDQ256) Assembly() string {
	return assemblyVpmovsxdq256
}

func (v *VPMOVSXDQ256) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [4]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxdq256(&vals, &ret)

	log.Printf("VPMOVSXDQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
