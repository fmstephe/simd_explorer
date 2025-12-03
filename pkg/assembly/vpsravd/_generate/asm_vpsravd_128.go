package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsravd_128.go -out ../asm_vpsravd_128.s -stubs ../stub_vpsravd_128.go -pkg vpsravd
func main() {
	TEXT("vpsravd128", NOSPLIT, "func(vals *[4]int32, shifts *[4]uint32, ret *[4]int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := XMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSRAVD: arithmetic right shift packed doubleword integers by variable counts")
	VPSRAVD(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
