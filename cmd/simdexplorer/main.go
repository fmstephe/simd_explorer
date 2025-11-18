package main

import (
	"log"
	"os"

	"github.com/fmstephe/simd_explorer/pkg/ui"
)

func main() {
	if f, err := os.OpenFile("./debug.log", os.O_TRUNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("Error opening log file %v", err)
	} else {
		log.SetOutput(f)
	}
	ui.Run()
}
