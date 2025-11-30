package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128.s
var assemblyVpermilps128 string

//go:embed stub_vpermilps_128.go
var stubVpermilps128 string

type VPERMILPS128 struct {
	vals    *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPERMILPS128() *VPERMILPS128 {
	return &VPERMILPS128{
		vals:    number.NewNamedFloatParameter("vals", 128, 32),
		control: number.NewNamedUintParameter("control", 128, 32, 16),
		ret:     number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals, v.control}
}

func (v *VPERMILPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128) Name() string {
	return "VPERMILPS (128 bit) reg-control"
}

func (v *VPERMILPS128) Description() string {
	return "Permute single-precision floats using per-lane 2-bit selectors from control register."
}

func (v *VPERMILPS128) Stub() string {
	return stubVpermilps128
}

func (v *VPERMILPS128) Assembly() string {
	return assemblyVpermilps128
}

func (v *VPERMILPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	control := [4]float32{}
	copy(control[:], number.ToFloat32Slice(v.control.FlatData()))
	ret := [4]float32{}

	vpermilps128(&vals, &control, &ret)

	log.Printf("VPERMILPS128 vals %v control %v ret %v", vals, control, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
