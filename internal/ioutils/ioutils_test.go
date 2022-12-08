package ioutils

import (
	"bytes"
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"testing"
)

// TestReadFileAsStringSlices tests the function readFileAsStringSlices.
func TestReadFileAsStringSlices(t *testing.T) {
	var buffer bytes.Buffer

	// Test 1: A file containing only one line.
	buffer.WriteString("Hello World")

	obtained, err := readFileAsStringSlices(&buffer)

	if err != nil {
		t.Errorf("Read file error: Unexpected error %v", err)
	}

	expected := []string{"Hello World"}

	if !sliceutils.EqualStringSlices(obtained, expected) {
		t.Errorf("Read file error: Expected %v but got %v", expected, obtained)
	}

	buffer.Reset()

	// Test 2: A file containing multiple lines.
	buffer.WriteString("Hello\nWorld\nGoodbye\nWorld")

	obtained, err = readFileAsStringSlices(&buffer)

	if err != nil {
		t.Errorf("Read file error: Unexpected error %v", err)
	}

	expected = []string{"Hello", "World", "Goodbye", "World"}

	if !sliceutils.EqualStringSlices(obtained, expected) {
		t.Errorf("Read file error: Expected %v but got %v", expected, obtained)
	}

}
