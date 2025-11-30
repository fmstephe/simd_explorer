package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_psadbw_128.s
var assemblyPsadbw128 string

//go:embed stub_psadbw_128.go
var stubPsadbw128 string

type PSADBW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPSADBW128() *PSADBW128 {
	return &PSADBW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *PSADBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PSADBW128) Output() *number.Parameter {
	return v.ret
}

func (v *PSADBW128) Name() string {
	return "PSADBW (128 bit)"
}

func (v *PSADBW128) Description() string {
	return "Packed sum of absolute byte differences; two 64-bit lane sums."
}

func (v *PSADBW128) Stub() string {
	return stubPsadbw128
}

func (v *PSADBW128) Assembly() string {
	return assemblyPsadbw128
}

func (v *PSADBW128) Run() (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [8]uint16{}

	psadbw128(&vals1, &vals2, &ret)

	log.Printf("PSADBW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *PSADBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
