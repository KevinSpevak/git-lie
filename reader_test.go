package main

import (
	"strings"
	"testing"
)

func TestStripLies(t *testing.T) {
	corpus := "this line should be included\n" +
		"this line should be excluded !git-lie\n" +
		"this line should also be included\n" +
		"the following content should be excluded" +
		"<git-lie>hello world" +
		"this won't appear" +
		"</git-lie>" +
		"but this one will"

	stripped := StripLies(corpus)

	if strings.Contains(stripped, "excluded") {
		t.FailNow()
	}

	if strings.Contains(stripped, "this won't appear") {
		t.FailNow()
	}
}
