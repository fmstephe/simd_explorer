package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovpd_256_load.go -out ../asm_vmaskmovpd_256_load.s -stubs ../stub_vmaskmovpd_256_load.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovpd256Load", NOSPLIT, "func(vals *[4]float64, mask *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load mask into YMM register")
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)
	Comment("Zero destination")
	regDst := YMM()
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
