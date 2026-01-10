package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vinserti128_256_zero.go -out ../asm_vinserti128_256_zero.s -stubs ../stub_vinserti128_256_zero.go -pkg vinserti128
func main() {
	TEXT("vinserti128256Zero", NOSPLIT, "func(vals128 *[4]uint32, vals256 *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals128 := Load(Param("vals128"), GP64())
	vals256 := Load(Param("vals256"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals128 into XMM register")
	vals128X := XMM()
	VMOVDQU(Mem{Base: vals128}, vals128X)
	Comment("Load vals256 into YMM register")
	vals256Y := YMM()
	VMOVDQU(Mem{Base: vals256}, vals256Y)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VINSERTI128(U8(0x00), vals128X, vals256Y, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
