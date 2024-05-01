package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	input := "field1\tfield2\tfield3\nfield4\tfield5\tfield6"
	expected := "field2\tfield3\nfield5\tfield6"

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Expected '%s'", err)
	}
	oldStdout := os.Stdout
	os.Stdout = w

	cut(strings.NewReader(input), "\t", []int{2, 3}, false)

	w.Close()
	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)
	result := strings.TrimSuffix(string(out), "\n")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
