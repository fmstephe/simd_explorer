package pmaxsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsw_256.s
var assemblyVpmaxsw256 string

//go:embed stub_vpmaxsw_256.go
var stubVpmaxsw256 string

type VPMAXSW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSW256() *VPMAXSW256 {
	return &VPMAXSW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMAXSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSW256) Name() string {
	return "VPMAXSW (256 bit)"
}

func (v *VPMAXSW256) Description() string {
	return "Packed max of signed 16-bit words per lane (VEX, per 128-bit lane)."
}

func (v *VPMAXSW256) Stub() string {
	return stubVpmaxsw256
}

func (v *VPMAXSW256) Assembly() string {
	return assemblyVpmaxsw256
}

func (v *VPMAXSW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [16]int16{}

	vpmaxsw256(&vals1, &vals2, &ret)

	log.Printf("VPMAXSW256 input %v %v output %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPMAXSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
