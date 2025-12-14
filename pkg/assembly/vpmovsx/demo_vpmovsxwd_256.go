package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxwd_256.s
var assemblyVpmovsxwd256 string

//go:embed stub_vpmovsxwd_256.go
var stubVpmovsxwd256 string

type VPMOVSXWD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXWD256() *VPMOVSXWD256 {
	return &VPMOVSXWD256{
		vals: number.NewNamedIntParameter("vals", 128, 16, 10),
		ret:  number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMOVSXWD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXWD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXWD256) Name() string {
	return "VPMOVSXWD (256 bit) "
}

func (v *VPMOVSXWD256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVSXWD256) Stub() string {
	return stubVpmovsxwd256
}

func (v *VPMOVSXWD256) Assembly() string {
	return assemblyVpmovsxwd256
}

func (v *VPMOVSXWD256) Run() {
	vals := [8]int16{}
	copy(vals[:], number.ToInt16Slice(v.vals.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmovsxwd256(&vals, &ret)

	log.Printf("VPMOVSXWD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXWD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
