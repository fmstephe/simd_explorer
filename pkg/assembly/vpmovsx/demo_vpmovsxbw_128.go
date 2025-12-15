package vpmovsx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovsxbw_128.s
var assemblyVpmovsxbw128 string

//go:embed stub_vpmovsxbw_128.go
var stubVpmovsxbw128 string

type VPMOVSXBW128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVSXBW128() *VPMOVSXBW128 {
	return &VPMOVSXBW128{
		vals: number.NewNamedIntParameter("vals", 128, 8, 10),
		ret:  number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMOVSXBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVSXBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVSXBW128) Name() string {
	return "VPMOVSXBW (128 bit) "
}

func (v *VPMOVSXBW128) Description() string {
	return "Sign-extend packed 8-bit integers to 16-bit integers."
}

func (v *VPMOVSXBW128) Stub() string {
	return stubVpmovsxbw128
}

func (v *VPMOVSXBW128) Assembly() string {
	return assemblyVpmovsxbw128
}

func (v *VPMOVSXBW128) Run() {
	vals := [16]int8{}
	copy(vals[:], number.ToInt8Slice(v.vals.FlatData()))
	ret := [8]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpmovsxbw128(&vals, &ret)

	log.Printf("VPMOVSXBW vals %v ret %v", vals, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVSXBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
