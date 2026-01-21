package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fmstephe/simd_explorer/pkg/ui"
)

func main() {
	f, err := truncateAndOpenLog()
	if err != nil {
		log.Fatalf("Error preparing log file %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	ui.Run()
}

func truncateAndOpenLog() (*os.File, error) {
	logPath := filepath.Join(os.TempDir(), "simdexplorer.log")

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}

	return os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}
