package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandnps_256.s
var assemblyVandnps256 string

//go:embed stub_vandnps_256.go
var stubVandnps256 string

type VANDNPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVANDNPS256() *VANDNPS256 {
	return &VANDNPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VANDNPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VANDNPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VANDNPS256) Name() string {
	return "VANDNPS (256 bit)"
}

func (v *VANDNPS256) Description() string {
	return "Bitwise AND NOT with VEX encoding (per 128-bit lane); output shown as 32-bit hex lanes."
}

func (v *VANDNPS256) Stub() string {
	return stubVandnps256
}

func (v *VANDNPS256) Assembly() string {
	return assemblyVandnps256
}

func (v *VANDNPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vandnps256(&vals1, &vals2, &ret)

	log.Printf("VANDNPS256 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VANDNPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
