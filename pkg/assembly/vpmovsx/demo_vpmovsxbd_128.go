package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbd_128.s
var assemblyVpmovsxbd128 string

//go:embed stub_vpmovsxbd_128.go
var stubVpmovsxbd128 string

type VPMOVSXBD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBD128() *VPMOVSXBD128 {
	return &VPMOVSXBD128{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMOVSXBD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBD128) Name() string {
	return "VPMOVSXBD (128 bit) "
}

func (v *VPMOVSXBD128) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVSXBD128) Stub() string {
	return stubVpmovsxbd128
}

func (v *VPMOVSXBD128) Assembly() string {
	return assemblyVpmovsxbd128
}

func (v *VPMOVSXBD128) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmovsxbd128(&vals, &ret)

	log.Printf("VPMOVSXBD vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
