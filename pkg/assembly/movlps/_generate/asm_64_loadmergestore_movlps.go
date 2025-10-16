package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_64_loadmergestore_movlps.go -out ../asm_64_loadmergestore_movlps.s -stubs ../stub_64_loadmergestore_movlps.go -pkg movlps
func main() {
	TEXT("movlps64LoadMergeStoreMovlps", NOSPLIT, "func(lower, upper *[2]float32, ret *[4]float32)")

	Comment("load params")
	lower := Load(Param("lower"), GP64())
	upper := Load(Param("upper"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load upper (64 bit) into upper half of XMM register")
	regUpper := XMM()
	MOVHPS(Mem{Base: upper}, regUpper)

	Comment("Load lower (64 bit) and merge with upper in register")
	MOVLPS(Mem{Base: lower}, regUpper)

	Comment("Write contents of the merged XMM register into memory region")
	MOVUPS(regUpper, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
