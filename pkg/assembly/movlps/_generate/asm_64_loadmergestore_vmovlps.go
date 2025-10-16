package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_64_loadmergestore_vmovlps.go -out ../asm_64_loadmergestore_vmovlps.s -stubs ../stub_64_loadmergestore_vmovlps.go -pkg movlps
func main() {
	TEXT("movlps64LoadMergeStoreVmovlps", NOSPLIT, "func(lower, upper *[2]float32, ret *[4]float32)")

	Comment("load params")
	lower := Load(Param("lower"), GP64())
	upper := Load(Param("upper"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load upper (64 bit) into upper half of XMM register")
	regUpper := XMM()
	MOVHPS(Mem{Base: upper}, regUpper)

	Comment("Load lower (64 bit) and merge with upper in register")
	regMerged := XMM()
	VMOVLPS(Mem{Base: lower}, regUpper, regMerged)

	Comment("Write contents of the merged XMM register into memory region")
	VMOVUPS(regMerged, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
