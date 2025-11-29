package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_256.s
var assemblyVpermilpd256 string

//go:embed stub_vpermilpd_256.go
var stubVpermilpd256 string

type VPERMILPD256 struct {
	vals    *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPERMILPD256() *VPERMILPD256 {
	return &VPERMILPD256{
		vals:    number.NewNamedFloatParameter("vals", 256, 64),
		control: number.NewNamedUintParameter("control", 256, 64, 16),
		ret:     number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VPERMILPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals, v.control}
}

func (v *VPERMILPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPD256) Name() string {
	return "VPERMILPD (256 bit) reg-control"
}

func (v *VPERMILPD256) Description() string {
	return "Permute double-precision floats using per-lane 2-bit selectors from control register (per 128-bit lane)."
}

func (v *VPERMILPD256) Stub() string {
	return stubVpermilpd256
}

func (v *VPERMILPD256) Assembly() string {
	return assemblyVpermilpd256
}

func (v *VPERMILPD256) Run(_ [][]byte) (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	control := [4]float64{}
	copy(control[:], number.ToFloat64Slice(v.control.FlatData()))
	ret := [4]float64{}

	vpermilpd256(&vals, &control, &ret)

	log.Printf("VPERMILPD256 vals %v control %v ret %v", vals, control, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
