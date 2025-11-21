package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256.s
var assemblyVpermilps256 string

//go:embed stub_vpermilps_256.go
var stubVpermilps256 string

type VPERMILPS256 struct {
}

func (v *VPERMILPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),    // vals
		number.NewUintParameter(256, 32, 16), // control
	}
}

func (v *VPERMILPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256) Name() string {
	return "VPERMILPS (256 bit) reg-control"
}

func (v *VPERMILPS256) Description() string {
	return "Permute single-precision floats using per-lane 2-bit selectors from control register (per 128-bit lane)."
}

func (v *VPERMILPS256) Stub() string {
	return stubVpermilps256
}

func (v *VPERMILPS256) Assembly() string {
	return assemblyVpermilps256
}

func (v *VPERMILPS256) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))
	control := [8]float32{}
	copy(control[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vpermilps256(&vals, &control, &ret)

	log.Printf("VPERMILPS256 vals %v control %v ret %v", vals, control, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
