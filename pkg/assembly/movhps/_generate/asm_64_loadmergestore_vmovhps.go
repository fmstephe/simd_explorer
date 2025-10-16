package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_64_loadmergestore_vmovhps.go -out ../asm_64_loadmergestore_vmovhps.s -stubs ../stub_64_loadmergestore_vmovhps.go -pkg movhps
func main() {
	TEXT("movhps64LoadMergeStoreVmovhps", NOSPLIT, "func(lower, upper *[2]float32, ret *[4]float32)")

	Comment("load params")
	lower := Load(Param("lower"), GP64())
	upper := Load(Param("upper"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load lower (64 bit) into lower half of XMM register")
	regLower := XMM()
	MOVLPS(Mem{Base: lower}, regLower)

	Comment("Load upper (64 bit) and merge with lower in register")
	regMerged := XMM()
	VMOVHPS(Mem{Base: upper}, regLower, regMerged)

	Comment("Write contents of the merged XMM register into memory region")
	VMOVUPS(regMerged, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
