package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movaps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movups"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqa"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastb"
)

func Instructions() []assembly.Instruction {
	return []assembly.Instruction{
		&vpbroadcastb.Vpbroadcastb128{},
		&vpbroadcastb.Vpbroadcastb256{},
		&vpbroadcastb.Vpbroadcastb512{},
		&vpbroadcastb.Vpbroadcastb128K{},
		&vpbroadcastb.Vpbroadcastb256K{},
		&vpbroadcastb.Vpbroadcastb512K{},
		&vmovdqu.VMOVDQU128LoadStore{},
		&vmovdqu.VMOVDQU256LoadStore{},
		&vmovdqa.VMOVDQA128LoadStore{},
		&vmovdqa.VMOVDQA256LoadStore{},
		&addps.ADDPS128{},
		&addps.VADDPS128{},
		&addps.VADDPS256{},
		&movss.MOVSS128LoadStore{},
		&movss.VMOVSS128LoadStore{},
		&movaps.MOVAPS128LoadStore{},
		&movaps.VMOVAPS128LoadStore{},
		&movaps.VMOVAPS256LoadStore{},
		&movups.MOVUPS128LoadStore{},
		&movups.VMOVUPS128LoadStore{},
		&movups.VMOVUPS256LoadStore{},
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
