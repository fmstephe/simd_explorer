package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubd_128.s
var assemblyVpsubd128 string

//go:embed stub_vpsubd_128.go
var stubVpsubd128 string

type VPSUBD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBD128() *VPSUBD128 {
	return &VPSUBD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSUBD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBD128) Name() string {
	return "VPSUBD (128 bit) "
}

func (v *VPSUBD128) Description() string {
	return "Subtract packed u32 doublewords (wrap-around)."
}

func (v *VPSUBD128) Stub() string {
	return stubVpsubd128
}

func (v *VPSUBD128) Assembly() string {
	return assemblyVpsubd128
}

func (v *VPSUBD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpsubd128(&vals1, &vals2, &ret)

	log.Printf("VPSUBD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSUBD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
