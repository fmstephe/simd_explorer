package ui

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly/all"
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
	commandSearch := commands.NewCommandSearch(all.Instructions(), app)

	app.Push(commandSearch.GetBox())

	if err := app.Run(); err != nil {
		panic(err)
	}
}
