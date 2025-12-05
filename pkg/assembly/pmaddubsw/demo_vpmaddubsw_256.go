package pmaddubsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddubsw_256.s
var assemblyVpmaddubsw256 string

//go:embed stub_vpmaddubsw_256.go
var stubVpmaddubsw256 string

type VPMADDUBSW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDUBSW256() *VPMADDUBSW256 {
	return &VPMADDUBSW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMADDUBSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDUBSW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDUBSW256) Name() string {
	return "VPMADDUBSW (256 bit) "
}

func (v *VPMADDUBSW256) Description() string {
	return "Multiply unsigned bytes by signed bytes, add adjacent products, saturate to signed 16-bit."
}

func (v *VPMADDUBSW256) Stub() string {
	return stubVpmaddubsw256
}

func (v *VPMADDUBSW256) Assembly() string {
	return assemblyVpmaddubsw256
}

func (v *VPMADDUBSW256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]int16{}

	vpmaddubsw256(&vals1, &vals2, &ret)

	log.Printf("VPMADDUBSW256 vals1 %v vals2(bytes) %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMADDUBSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
