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
	vals    *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPERMILPD128() *VPERMILPD128 {
	return &VPERMILPD128{
		vals:    number.NewNamedFloatParameter("vals", 128, 64),
		control: number.NewNamedUintParameter("control", 128, 64, 16),
		ret:     number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VPERMILPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.control,
	}
}

func (v *VPERMILPD128) Output() *number.Parameter {
	return v.ret
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

func (v *VPERMILPD128) Run() {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	control := [2]float64{}
	copy(control[:], number.ToFloat64Slice(v.control.FlatData()))

	ret := [2]float64{}

	vpermilpd128(&vals, &control, &ret)

	log.Printf("VPERMILPD128 vals %v control %v ret %v", vals, control, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
