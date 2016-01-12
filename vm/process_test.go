package vm

import (
	"testing"
)

func TestProcess(t *testing.T) {
	parsed, err := Parse([]string{
		"LOAD A,1",
		"PUSH A",
		"POP A",
		"PRINT A",
		"PRINTLN",
		"END",
	}, NewOption())
	if err != nil {
		t.Fatalf("parseing failed: \"%s\"", err.Error())
	}
	result, err := Process(parsed, NewOption())
	if err != nil {
		t.Fatalf("processing failed: \"%s\"", err.Error())
	}
	expected := "1\n"
	if result != expected {
		t.Errorf("\"%s\" <> \"%s\"", result, expected)
	}
}
