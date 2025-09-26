package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
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
