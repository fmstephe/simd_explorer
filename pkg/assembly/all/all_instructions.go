package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cmpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/maxps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/maxss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/minps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/minss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movaps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movups"
	"github.com/fmstephe/simd_explorer/pkg/assembly/mulps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/mulss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqa"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastb"
)

func Instructions() []assembly.Instruction {
	return []assembly.Instruction{
		// vpbroadcastb
		&vpbroadcastb.VPBROADCASTB128{},
		&vpbroadcastb.VPBROADCASTB256{},
		&vpbroadcastb.VPBROADCASTB512{},
		&vpbroadcastb.VPBROADCASTB128K{},
		&vpbroadcastb.VPBROADCASTB256K{},
		&vpbroadcastb.VPBROADCASTB512K{},
		// vmovdqu
		&vmovdqu.VMOVDQU128{},
		&vmovdqu.VMOVDQU256{},
		&vmovdqa.VMOVDQA128{},
		&vmovdqa.VMOVDQA256{},
		// addps
		&addps.ADDPS128{},
		&addps.VADDPS128{},
		&addps.VADDPS256{},
		// mulps
		&mulps.MULPS128{},
		&mulps.VMULPS128{},
		&mulps.VMULPS256{},
		// subps
		&subps.SUBPS128{},
		&subps.VSUBPS128{},
		&subps.VSUBPS256{},
		// addss
		&addss.ADDSS128{},
		&addss.VADDSS128{},
		// divss
		&divss.DIVSS128{},
		&divss.VDIVSS128{},
		// divps
		&divps.DIVPS128{},
		&divps.VDIVPS128{},
		&divps.VDIVPS256{},
		// subss
		&subss.SUBSS128{},
		&subss.VSUBSS128{},
		// mulss
		&mulss.MULSS128{},
		&mulss.VMULSS128{},
		// movss
		&movss.MOVSS128{},
		&movss.VMOVSS128{},
		// cmpss
		&cmpss.CMPSS128EQ{},
		&cmpss.CMPSS128LT{},
		&cmpss.CMPSS128LE{},
		&cmpss.CMPSS128UNORD{},
		&cmpss.CMPSS128NEQ{},
		&cmpss.CMPSS128NLT{},
		&cmpss.CMPSS128NLE{},
		&cmpss.CMPSS128ORD{},
		&cmpss.VCMPSS128EQ{},
		&cmpss.VCMPSS128LT{},
		&cmpss.VCMPSS128LE{},
		&cmpss.VCMPSS128UNORD{},
		&cmpss.VCMPSS128NEQ{},
		&cmpss.VCMPSS128NLT{},
		&cmpss.VCMPSS128NLE{},
		&cmpss.VCMPSS128ORD{},
		// rcpps
		&rcpps.RCPPS128{},
		&rcpps.VRCPPS128{},
		&rcpps.VRCPPS256{},
		// sqrtps
		&sqrtps.SQRTPS128{},
		&sqrtps.VSQRTPS128{},
		&sqrtps.VSQRTPS256{},
		// rsqrtps
		&rsqrtps.RSQRTPS128{},
		&rsqrtps.VRSQRTPS128{},
		&rsqrtps.VRSQRTPS256{},
		// rcpss
		&rcpss.RCPSS128{},
		&rcpss.VRCPSS128{},
		// rsqrtss
		&rsqrtss.RSQRTSS128{},
		&rsqrtss.VRSQRTSS128{},
		// sqrtss
		&sqrtss.SQRTSS128{},
		&sqrtss.VSQRTSS128{},
		// movaps
		&movaps.MOVAPS128{},
		&movaps.VMOVAPS128{},
		&movaps.VMOVAPS256{},
		// movups
		&movups.MOVUPS128{},
		&movups.VMOVUPS128{},
		&movups.VMOVUPS256{},
		// movlps
		&movlps.MOVLPS64{},
		&movlps.VMOVLPS64{},
		// movhpos
		&movhps.MOVHPS64{},
		&movhps.VMOVHPS64{},
		// movlhps
		&movlhps.MOVLHPS64{},
		&movlhps.VMOVLHPS64{},
		// maxps
		&maxps.MAXPS128{},
		&maxps.VMAXPS128{},
		&maxps.VMAXPS256{},
		// maxss
		&maxss.MAXSS128{},
		&maxss.VMAXSS128{},
		// minps
		&minps.MINPS128{},
		&minps.VMINPS128{},
		&minps.VMINPS256{},
		// minss
		&minss.MINSS128{},
		&minss.VMINSS128{},
		// movmskps
		&movmskps.MOVMSKPS128{},
		&movmskps.VMOVMSKPS128{},
		&movmskps.VMOVMSKPS256{},
		// movmskpd
		&movmskpd.MOVMSKPD128{},
		&movmskpd.VMOVMSKPD128{},
		&movmskpd.VMOVMSKPD256{},
	}
}

func SupportedInstructions() []assembly.Instruction {
	instructions := Instructions()
	supported := make([]assembly.Instruction, 0, len(instructions))
	for _, inst := range instructions {
		if inst.Supported() {
			supported = append(supported, inst)
		}
	}
	return supported
}
