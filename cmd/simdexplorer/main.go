package main

import (
	"log"
	"os"

	"github.com/fmstephe/simd_explorer/pkg/ui"
)

func main() {
	if f, err := os.OpenFile("./simd_explorer.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("Error opening log file %v", err)
	} else {
		log.SetOutput(f)
	}
	ui.Run()
}
