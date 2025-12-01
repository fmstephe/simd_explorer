package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermd_256.go -out ../asm_vpermd_256.s -stubs ../stub_vpermd_256.go -pkg vpermd
func main() {
	TEXT("vpermd256", NOSPLIT, "func(vals *[8]uint32, control *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load control and source into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regCtrl := YMM()
	VMOVDQA(Mem{Base: control}, regCtrl)

	Comment("VPERMD: permute packed 32-bit integers with per-lane dword indices")
	VPERMD(regVals, regCtrl, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
