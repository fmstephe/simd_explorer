package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsadbw_256.s
var assemblyVpsadbw256 string

//go:embed stub_vpsadbw_256.go
var stubVpsadbw256 string

type VPSADBW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSADBW256() *VPSADBW256 {
	return &VPSADBW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSADBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSADBW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSADBW256) Name() string {
	return "VPSADBW (256 bit)"
}

func (v *VPSADBW256) Description() string {
	return "Packed sum of absolute byte differences (VEX); four 64-bit lane sums."
}

func (v *VPSADBW256) Stub() string {
	return stubVpsadbw256
}

func (v *VPSADBW256) Assembly() string {
	return assemblyVpsadbw256
}

func (v *VPSADBW256) Run() (output []byte) {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint16{}

	vpsadbw256(&vals1, &vals2, &ret)

	log.Printf("VPSADBW256 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPSADBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
