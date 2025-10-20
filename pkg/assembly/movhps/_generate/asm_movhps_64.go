package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_movhps_64.go -out ../asm_movhps_64.s -stubs ../stub_movhps_64.go -pkg movhps
func main() {
	TEXT("movhps64", NOSPLIT, "func(lower, upper *[2]float32, ret *[4]float32)")
	Comment("load params")
	lower := Load(Param("lower"), GP64())
	upper := Load(Param("upper"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load lower (64 bit) into lower half of XMM register")
	regLower := XMM()
	MOVLPS(Mem{Base: lower}, regLower)

	Comment("Load upper (64 bit) and merge with lower in register")
	MOVHPS(Mem{Base: upper}, regLower)

	Comment("Write contents of the merged XMM register into memory region")
	VMOVUPS(regLower, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}