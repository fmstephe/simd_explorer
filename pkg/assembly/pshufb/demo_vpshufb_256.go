package pshufb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufb_256.s
var assemblyVpshufb256 string

//go:embed stub_vpshufb_256.go
var stubVpshufb256 string

type VPSHUFB256 struct {
	vals1   *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPSHUFB256() *VPSHUFB256 {
	return &VPSHUFB256{
		vals1:   number.NewNamedUintParameter("vals1", 256, 8, 10),
		control: number.NewNamedUintParameter("control", 256, 8, 10),
		ret:     number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSHUFB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.control,
	}
}

func (v *VPSHUFB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFB256) Name() string {
	return "VPSHUFB (256 bit) "
}

func (v *VPSHUFB256) Description() string {
	return "Shuffle bytes in vals1 according to control (per 128-bit lane); high bit zeroes out."
}

func (v *VPSHUFB256) Stub() string {
	return stubVpshufb256
}

func (v *VPSHUFB256) Assembly() string {
	return assemblyVpshufb256
}

func (v *VPSHUFB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], number.ToUint8Slice(v.vals1.FlatData()))
	control := [32]uint8{}
	copy(control[:], number.ToUint8Slice(v.control.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpshufb256(&vals1, &control, &ret)

	log.Printf("VPSHUFB256 vals1 %v control %v ret %v", vals1, control, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSHUFB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
