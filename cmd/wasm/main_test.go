package main

import (
	"testing"
)

func TestFormatShellScript(t *testing.T) {
	formatted, err := formatShellScript("  hello!  \n  New line!  \n")
	if err != nil {
		t.Errorf("Returned unexpected error: %s", err)
	}
	if formatted != "hello!\nNew line!\n" {
		t.Errorf("Did not correctly format. Got = %s", formatted)
	}
}
