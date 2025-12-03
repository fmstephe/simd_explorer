package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsravd_256.go -out ../asm_vpsravd_256.s -stubs ../stub_vpsravd_256.go -pkg vpsravd
func main() {
	TEXT("vpsravd256", NOSPLIT, "func(vals *[8]int32, shifts *[8]uint32, ret *[8]int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := YMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSRAVD: arithmetic right shift packed doubleword integers by variable counts")
	VPSRAVD(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
