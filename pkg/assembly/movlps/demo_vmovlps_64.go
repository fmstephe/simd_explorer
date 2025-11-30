package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovlps_64.s
var assemblyVmovlps64 string

//go:embed stub_vmovlps_64.go
var stubVmovlps64 string

type VMOVLPS64 struct {
	lower *number.Parameter
	upper *number.Parameter
	ret   *number.Parameter
}

func NewVMOVLPS64() *VMOVLPS64 {
	return &VMOVLPS64{
		lower: number.NewNamedFloatParameter("lower", 64, 32),
		upper: number.NewNamedFloatParameter("upper", 64, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMOVLPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.lower,
		v.upper,
	}
}

func (v *VMOVLPS64) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVLPS64) Name() string {
	return "VMOVLPS XMM (2X 64 bit)"
}

func (v *VMOVLPS64) Description() string {
	return "AVX form: move two floats into the low 64 bits of XMM; high 64 supplied separately."
}

func (v *VMOVLPS64) Stub() string {
	return stubVmovlps64
}

func (v *VMOVLPS64) Assembly() string {
	return assemblyVmovlps64
}

func (v *VMOVLPS64) Run() (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(v.lower.FlatData()))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(v.upper.FlatData()))

	ret := [4]float32{}

	vmovlps64(&lower, &upper, &ret)

	log.Printf("VMOVLPS64 input lower %v upper %v output %v", lower, upper, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVLPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
