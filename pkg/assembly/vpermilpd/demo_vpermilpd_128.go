package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_128.s
var assemblyVpermilpd128 string

//go:embed stub_vpermilpd_128.go
var stubVpermilpd128 string

type VPERMILPD128 struct {
}

func (v *VPERMILPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),    // vals
		number.NewUintParameter(128, 64, 16), // control
	}
}

func (v *VPERMILPD128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64)
}

func (v *VPERMILPD128) Name() string {
	return "VPERMILPD (128 bit) reg-control"
}

func (v *VPERMILPD128) Description() string {
	return "Permute double-precision floats using per-lane 2-bit selectors from control register."
}

func (v *VPERMILPD128) Stub() string {
	return stubVpermilpd128
}

func (v *VPERMILPD128) Assembly() string {
	return assemblyVpermilpd128
}

func (v *VPERMILPD128) Run(inputs [][]byte) (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(inputs[0]))
	control := [2]float64{}
	copy(control[:], number.ToFloat64Slice(inputs[1]))

	ret := [2]float64{}

	vpermilpd128(&vals, &control, &ret)

	log.Printf("VPERMILPD128 vals %v control %v ret %v", vals, control, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VPERMILPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
