package main

import (
	"net/http"
	"strings"
)

// lineEndsBlockLie returns true if the line contains an inline lie marker
func lineIsInlineLie(line string, marker string) bool {
	return strings.Contains(line, marker)
}

// lineEndsBlockLie returns true if the line is the beginning of a block lie
func lineBeginsBlockLie(line string, marker string) bool {
	return strings.Contains(line, marker)
}

// lineEndsBlockLie returns true if the line is the end of a block lie
func lineEndsBlockLie(line string, marker string) bool {
	return strings.Contains(line, marker)
}

// IsTextFile returns true iff the FILE argument has a MIME type starting with "text/".
func IsTextFile(file string) bool {
	return (strings.HasPrefix(http.DetectContentType([]byte(file)), "text/"))
}

// StripLies returns a copy of the input file's bytes with lies removed
func StripLies(file string) string {

	lines := strings.Split(file, string("\n"))

	// We know the processed file will have at most the same number of lines
	output := make([]string, len(lines))

	// true if the loop is inside a block lie
	inBlock := false

	// index into the list of output lines
	outputIdx := 0

	for _, line := range lines {
		if lineBeginsBlockLie(line, "<git-lie>") {
			inBlock = true
		}

		if !inBlock && !lineIsInlineLie(line, "!git-lie") {
			output[outputIdx] = line
			outputIdx += 1
		}

		// order is important - this way, lines containing </git-lie> are excluded
		if lineEndsBlockLie(line, "</git-lie>") {
			inBlock = false
		}

	}

	return strings.Join(output, string("\n"))
}
