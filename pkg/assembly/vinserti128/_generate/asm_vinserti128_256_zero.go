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

	Comment("Load vals256 into YMM register")
	regY := YMM()
	VMOVDQU(Mem{Base: vals256}, regY)

	Comment("Insert 128-bit block into lower 128-bit lane (0) of YMM; upper lane preserved from vals256")
	VINSERTI128(U8(0x00), Mem{Base: vals128}, regY, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}


