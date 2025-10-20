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
}

func (v *VMOVHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *VMOVHPS64) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVHPS64) Name() string {
	return "VMOVHPS (2X 64 bit)"
}

func (v *VMOVHPS64) Description() string {
	return "TODO"
}

func (v *VMOVHPS64) Stub() string {
	return stubVmovhps64
}

func (v *VMOVHPS64) Assembly() string {
	return assemblyVmovhps64
}

func (v *VMOVHPS64) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmovhps64(&lower, &upper, &ret)

	log.Printf("VMOVHPS64 input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
