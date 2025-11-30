package movlhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovlhps_64.s
var assemblyVmovlhps64 string

//go:embed stub_vmovlhps_64.go
var stubVmovlhps64 string

type VMOVLHPS64 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVLHPS64() *VMOVLHPS64 {
	return &VMOVLHPS64{
		vals: number.NewNamedFloatParameter("vals", 64, 32),
		ret:  number.NewNamedFloatParameter("ret", 64, 32),
	}
}

func (v *VMOVLHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVLHPS64) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVLHPS64) Name() string {
	return "VMOVLHPS (64 bit) "
}

func (v *VMOVLHPS64) Description() string {
	return "AVX form: move low 64 bits of source into high 64 of destination XMM; low 64 preserved."
}

func (v *VMOVLHPS64) Stub() string {
	return stubVmovlhps64
}

func (v *VMOVLHPS64) Assembly() string {
	return assemblyVmovlhps64
}

func (v *VMOVLHPS64) Run() (output []byte) {
	vals := [2]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [2]float32{}

	vmovlhps64(&vals, &ret)

	log.Printf("VMOVLHPS64 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVLHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
