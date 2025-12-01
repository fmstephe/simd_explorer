package vpermps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermps_256.s
var assemblyVpermps256 string

//go:embed stub_vpermps_256.go
var stubVpermps256 string

type VPERMPS256 struct {
	vals    *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPERMPS256() *VPERMPS256 {
	return &VPERMPS256{
		vals:    number.NewNamedFloatParameter("vals", 256, 32),
		control: number.NewNamedUintParameter("control", 256, 32, 10),
		ret:     number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERMPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.control,
	}
}

func (v *VPERMPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMPS256) Name() string {
	return "VPERMPS (256 bit) "
}

func (v *VPERMPS256) Description() string {
	return "Permute 8 f32 elements from vals using per-dword indices in control (AVX2 VPERMPS; indices select within 128-bit lanes)."
}

func (v *VPERMPS256) Stub() string {
	return stubVpermps256
}

func (v *VPERMPS256) Assembly() string {
	return assemblyVpermps256
}

func (v *VPERMPS256) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	control := [8]uint32{}
	copy(control[:], number.ToUint32Slice(v.control.FlatData()))

	ret := [8]float32{}

	vpermps256(&vals, &control, &ret)

	log.Printf("VPERMPS256 control %v vals %v ret %v", control, vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
