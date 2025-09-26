package asmutil

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/sys/cpu"
)

var requiresRegex = regexp.MustCompile("^// Requires: (.*)$")

func IsSupported(assembly string) bool {
	lines := strings.Split(assembly, "\n")
	for _, line := range lines {
		features := getFeatures(line)
		for _, feature := range features {
			if !hasFeature(feature) {
				return false
			}
		}
	}

	return true
}

// Example features line "// Requires: AVX, AVX2, AVX512F, AVX512VL, SSE2"
func getFeatures(line string) []string {
	features := []string{}
	matches := requiresRegex.FindStringSubmatch(line)
	if len(matches) > 1 {
		allFeatures := matches[1]
		features = append(features, strings.Split(allFeatures, ", ")...)
	}

	return features
}

func hasFeature(feature string) bool {
	switch feature {
	case "SSE2":
		return cpu.X86.HasSSE2
	case "AVX":
		return cpu.X86.HasAVX
	case "AVX2":
		return cpu.X86.HasAVX2
	case "AVX512F":
		return cpu.X86.HasAVX512F
	case "AVX512VL":
		return cpu.X86.HasAVX512VL
	case "AVX512BW":
		return cpu.X86.HasAVX512BW
	default:
		panic(fmt.Errorf("Unknown feature name %q", feature))
	}
}
