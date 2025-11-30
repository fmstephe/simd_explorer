package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsadbw_128.s
var assemblyVpsadbw128 string

//go:embed stub_vpsadbw_128.go
var stubVpsadbw128 string

type VPSADBW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSADBW128() *VPSADBW128 {
	return &VPSADBW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSADBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSADBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSADBW128) Name() string {
	return "VPSADBW (128 bit)"
}

func (v *VPSADBW128) Description() string {
	return "Packed sum of absolute byte differences (VEX); two 64-bit lane sums."
}

func (v *VPSADBW128) Stub() string {
	return stubVpsadbw128
}

func (v *VPSADBW128) Assembly() string {
	return assemblyVpsadbw128
}

func (v *VPSADBW128) Run() (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [8]uint16{}

	vpsadbw128(&vals1, &vals2, &ret)

	log.Printf("VPSADBW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPSADBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
