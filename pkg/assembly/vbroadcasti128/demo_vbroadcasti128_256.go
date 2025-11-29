package vbroadcasti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vbroadcasti128_256.s
var assemblyVbroadcasti128256 string

//go:embed stub_vbroadcasti128_256.go
var stubVbroadcasti128256 string

type VBROADCASTI128256 struct {
	block *number.Parameter
	ret   *number.Parameter
}

func NewVBROADCASTI128256() *VBROADCASTI128256 {
	return &VBROADCASTI128256{
		block: number.NewNamedUintParameter("block", 128, 64, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VBROADCASTI128256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.block,
	}
}

func (v *VBROADCASTI128256) Output() *number.Parameter {
	return v.ret
}

func (v *VBROADCASTI128256) Name() string {
	return "VBROADCASTI128 (256 bit)"
}

func (v *VBROADCASTI128256) Description() string {
	return "Broadcast 128-bit block from memory to both 128-bit lanes of YMM."
}

func (v *VBROADCASTI128256) Stub() string {
	return stubVbroadcasti128256
}

func (v *VBROADCASTI128256) Assembly() string {
	return assemblyVbroadcasti128256
}

func (v *VBROADCASTI128256) Run(_ [][]byte) (output []byte) {
	var val [2]uint64
	copy(val[:], number.ToUint64Slice(v.block.FlatData()))
	var ret [4]uint64
	vbroadcasti128256(&val, &ret)
	log.Printf("VBROADCASTI128256 block %v output %v", val, ret)
	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VBROADCASTI128256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
