package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtpd2ps_256.go -out ../asm_vcvtpd2ps_256.s -stubs ../stub_vcvtpd2ps_256.go -pkg cvtpd2ps
func main() {
	TEXT("vcvtpd2ps256", NOSPLIT, "func(vals *[4]float64, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	valsY := YMM()
	VMOVDQU(Mem{Base: vals}, valsY)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VCVTPD2PSY(valsY, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
