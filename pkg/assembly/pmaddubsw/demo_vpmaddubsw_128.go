package pmaddubsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddubsw_128.s
var assemblyVpmaddubsw128 string

//go:embed stub_vpmaddubsw_128.go
var stubVpmaddubsw128 string

type VPMADDUBSW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDUBSW128() *VPMADDUBSW128 {
	return &VPMADDUBSW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMADDUBSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDUBSW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDUBSW128) Name() string {
	return "VPMADDUBSW (128 bit) "
}

func (v *VPMADDUBSW128) Description() string {
	return "Multiply unsigned bytes by signed bytes, add adjacent products, saturate to signed 16-bit."
}

func (v *VPMADDUBSW128) Stub() string {
	return stubVpmaddubsw128
}

func (v *VPMADDUBSW128) Assembly() string {
	return assemblyVpmaddubsw128
}

func (v *VPMADDUBSW128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [8]int16{}

	vpmaddubsw128(&vals1, &vals2, &ret)

	log.Printf("VPMADDUBSW128 vals1 %v vals2(bytes) %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMADDUBSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
