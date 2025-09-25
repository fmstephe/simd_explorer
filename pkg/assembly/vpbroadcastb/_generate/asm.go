package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm.go -out ../asm.s -stubs ../stub.go -pkg vpbroadcastb
func main() {
	vpbroadcast128()
	vpbroadcast256()
	vpbroadcast512()

	Generate()
}

func vpbroadcast128() {
	TEXT("vpbroadcastb128", NOSPLIT, "func(b byte, ret *[16]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into XMM register")
	regX := XMM()
	VPBROADCASTB(regXArg, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()
}

func vpbroadcast256() {
	TEXT("vpbroadcastb256", NOSPLIT, "func(b byte, ret *[32]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into YMM register")
	regY := YMM()
	VPBROADCASTB(regXArg, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()
}

func vpbroadcast512() {
	TEXT("vpbroadcastb512", NOSPLIT, "func(b byte, ret *[64]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into ZMM register")
	regZ := YMM()
	VPBROADCASTB(regXArg, regZ)

	Comment("Write contents of ZMM register into memory region")
	VMOVDQU64(regZ, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()
}
