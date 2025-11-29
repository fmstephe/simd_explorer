package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovups_256.s
var assemblyVmovups256 string

//go:embed stub_vmovups_256.go
var stubVmovups256 string

type VMOVUPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVUPS256() *VMOVUPS256 {
	return &VMOVUPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMOVUPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVUPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVUPS256) Name() string {
	return "VMOVUPS YMM (256 bit)"
}

func (v *VMOVUPS256) Description() string {
	return "Unaligned move of packed single-precision floats between memory and YMM; copies data unchanged."
}

func (v *VMOVUPS256) Stub() string {
	return stubVmovups256
}

func (v *VMOVUPS256) Assembly() string {
	return assemblyVmovups256
}

func (v *VMOVUPS256) Run(_ [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vmovups256(&vals, &ret)

	log.Printf("VMOVUPS256 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVUPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
