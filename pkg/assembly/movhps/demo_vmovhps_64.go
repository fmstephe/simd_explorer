package movhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovhps_64.s
var assemblyVmovhps64 string

//go:embed stub_vmovhps_64.go
var stubVmovhps64 string

type VMOVHPS64 struct {
	lower *number.Parameter
	upper *number.Parameter
	ret   *number.Parameter
}

func NewVMOVHPS64() *VMOVHPS64 {
	return &VMOVHPS64{
		lower: number.NewNamedFloatParameter("lower", 64, 32),
		upper: number.NewNamedFloatParameter("upper", 64, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMOVHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.lower,
		v.upper,
	}
}

func (v *VMOVHPS64) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVHPS64) Name() string {
	return "VMOVHPS (2X 64 bit)"
}

func (v *VMOVHPS64) Description() string {
	return "AVX form: move two floats into the high 64 bits of XMM; low 64 supplied separately."
}

func (v *VMOVHPS64) Stub() string {
	return stubVmovhps64
}

func (v *VMOVHPS64) Assembly() string {
	return assemblyVmovhps64
}

func (v *VMOVHPS64) Run(_ [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(v.lower.FlatData()))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(v.upper.FlatData()))

	ret := [4]float32{}

	vmovhps64(&lower, &upper, &ret)

	log.Printf("VMOVHPS64 input lower %v upper %v output %v", lower, upper, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
