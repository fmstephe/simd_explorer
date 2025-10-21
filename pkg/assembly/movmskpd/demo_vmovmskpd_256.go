package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskpd_256.s
var assemblyVmovmskpd256 string

//go:embed stub_vmovmskpd_256.go
var stubVmovmskpd256 string

type VMOVMSKPD256 struct {
}

func (v *VMOVMSKPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),
	}
}

func (v *VMOVMSKPD256) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VMOVMSKPD256) Name() string {
	return "VMOVMSKPD (256 bit) "
}

func (v *VMOVMSKPD256) Description() string {
	return "TODO"
}

func (v *VMOVMSKPD256) Stub() string {
	return stubVmovmskpd256
}

func (v *VMOVMSKPD256) Assembly() string {
	return assemblyVmovmskpd256
}

func (v *VMOVMSKPD256) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float64{}
	copy(floats[:], number.ToFloat64Slice(inputs[0]))

	ret := [4]byte{}

	vmovmskpd256(&floats, &ret)

	log.Printf("VMOVMSKPD256 input %v output %v", floats, ret)

	return ret[:]
}

func (v *VMOVMSKPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
