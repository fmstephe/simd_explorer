package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbq_128.s
var assemblyVpmovsxbq128 string

//go:embed stub_vpmovsxbq_128.go
var stubVpmovsxbq128 string

type VPMOVSXBQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBQ128() *VPMOVSXBQ128 {
	return &VPMOVSXBQ128{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVSXBQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBQ128) Name() string {
	return "VPMOVSXBQ (128 bit) "
}

func (v *VPMOVSXBQ128) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVSXBQ128) Stub() string {
	return stubVpmovsxbq128
}

func (v *VPMOVSXBQ128) Assembly() string {
	return assemblyVpmovsxbq128
}

func (v *VPMOVSXBQ128) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [2]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmovsxbq128(&vals, &ret)

	log.Printf("VPMOVSXBQ vals %v ret %v", vals, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
