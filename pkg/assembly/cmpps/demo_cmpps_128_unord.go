package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_unord.s
var assemblyCmpps128Unord string

//go:embed stub_cmpps_128_unord.go
var stubCmpps128Unord string

type CMPPS128UNORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128UNORD() *CMPPS128UNORD {
	return &CMPPS128UNORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128UNORD) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128UNORD) Name() string {
	return "CMPPS (128 bit) unord"
}

func (v *CMPPS128UNORD) Description() string {
	return "Compare packed single-precision floats for unordered (per lane)."
}

func (v *CMPPS128UNORD) Stub() string {
	return stubCmpps128Unord
}

func (v *CMPPS128UNORD) Assembly() string {
	return assemblyCmpps128Unord
}

func (v *CMPPS128UNORD) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Unord(&vals1, &vals2, &ret)

	log.Printf("CMPPS128UNORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPPS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
