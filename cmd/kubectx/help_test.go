package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_printHelp(t *testing.T) {
	var buf bytes.Buffer
	printHelp(&buf)

	out := buf.String()
	if !strings.Contains(out, "USAGE") {
		t.Errorf("Help string doesn't contain USAGE: ; output=%q", out)
	}

	if !strings.HasSuffix(out, "\n") {
		t.Errorf("Does not end with new line; output: %q", out)
	}
}
