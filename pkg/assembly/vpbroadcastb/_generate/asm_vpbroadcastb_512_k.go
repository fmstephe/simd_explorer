package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastb_512_k.go -out ../asm_vpbroadcastb_512_k.s -stubs ../stub_vpbroadcastb_512_k.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb512K", NOSPLIT, "func(b byte, k uint64, ret *[64]byte)")

	Comment("load params")
	b := Load(Param("b"), GP64())
	k := Load(Param("k"), K())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into YMM register")
	regZ := ZMM()
	VPBROADCASTB(regXArg, k, regZ)

	Comment("Write contents of ZMM register into memory region")
	VMOVDQU64(regZ, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}

