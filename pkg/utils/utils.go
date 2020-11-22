package utils

import (
	"bufio"
	"io"
)

func ReadLine(r io.Reader, lineNum int) (line string, err error) {
	b := bufio.NewScanner(r)
	var currentLine int

	for b.Scan() {
		currentLine++
		if currentLine == lineNum {
			return b.Text(), nil
		}
	}

	return line, io.EOF
}
