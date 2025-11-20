package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovpd_128_load.go -out ../asm_vmaskmovpd_128_load.s -stubs ../stub_vmaskmovpd_128_load.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovpd128Load", NOSPLIT, "func(vals *[2]float64, mask *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load mask into XMM register")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)
	Comment("Zero destination")
	regDst := XMM()
	VXORPS(regDst, regDst, regDst)

	Comment("Masked load: load selected double-precision elements from memory")
	VMASKMOVPD(Mem{Base: vals}, regMask, regDst)

	Comment("Write results into return memory address")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
