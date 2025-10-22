package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movaps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movups"
	"github.com/fmstephe/simd_explorer/pkg/assembly/mulss"
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
		// addss
		&addss.ADDSS128{},
		&addss.VADDSS128{},
		// divss
		&divss.DIVSS128{},
		&divss.VDIVSS128{},
		// subss
		&subss.SUBSS128{},
		&subss.VSUBSS128{},
		// mulss
		&mulss.MULSS128{},
		&mulss.VMULSS128{},
		// movss
		&movss.MOVSS128{},
		&movss.VMOVSS128{},
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
