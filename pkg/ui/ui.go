package ui

import (
	"github.com/fmstephe/simd_explorer/pkg/instructions"
	"github.com/fmstephe/simd_explorer/pkg/instructions/vpbroadcastb"
	"github.com/fmstephe/simd_explorer/pkg/ui/commands"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
)

func Run() {
	app := stackapp.NewStackApp()

	/*
		register256 := register.NewUIRegisterSet(app, 256)
		primitive := register256.Base2.GetPrimitive()
	*/
	// Setup the application with the components defined above
	commandSearch := commands.NewCommandSearch([]instructions.Instruction{&vpbroadcastb.VPBROADCASTB{}}, app)

	app.Push(commandSearch.GetBox())

	if err := app.Run(); err != nil {
		panic(err)
	}
}
