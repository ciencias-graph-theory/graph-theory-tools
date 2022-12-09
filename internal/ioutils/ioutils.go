// Package fileutils provides useful functions for I/O utilities.
package ioutils

import (
	"bufio"
	"io"
	"os"
)

// OpenFileAsStringSlices opens a file and returns its content as string slices.
func OpenFileAsStringSlices(filePath string) ([]string, error) {
	// Open the file.
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Read its contents.
	contents, err := readFileAsStringSlices(file)

	if err != nil {
		return nil, err
	}

	file.Close()

	// Return the contents.
	return contents, nil
}

// readFileAsStringSlices reads a file buffer and returns its content as a slice
// of strings. Each line corresponds to a string in the slice.
func readFileAsStringSlices(io io.Reader) ([]string, error) {
	// A slice of strings for each line in the file.
	var lines []string

	// Creates a scanner to read the file.
	fileScanner := bufio.NewScanner(io)

	// Split the file, line by line, using the scanner.
	fileScanner.Split(bufio.ScanLines)

	// Read each linea and append it to slice of strings.
	for fileScanner.Scan() {
		s := fileScanner.Text()
		lines = append(lines, s)
	}

	return lines, nil
}
