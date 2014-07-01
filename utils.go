package gohtml

import (
	"bytes"
	"strings"
)

// writeLine writes HTML to the buffer.
func writeRaw(bf *bytes.Buffer, s string) {
	bf.WriteString(s)
}

// writeLine writes an HTML line to the buffer.
func write(bf *bytes.Buffer, s string, indent int) {
	writeIndent(bf, indent)
	bf.WriteString(s)
}

// writeLine writes an HTML line to the buffer.
func writeLine(bf *bytes.Buffer, s string, indent int) {
	writeLineFeed(bf)
	writeIndent(bf, indent)
	bf.WriteString(s)
}

// writeLineFeed writes a line feed to the buffer.
func writeLineFeed(bf *bytes.Buffer) {
	if bf.Len() > 0 {
		bf.WriteString("\n")
	}
}

// writeIndent writes indents to the buffer.
func writeIndent(bf *bytes.Buffer, indent int) {
	bf.WriteString(strings.Repeat(defaultIndentString, indent))
}

// unifyLineFeed unifies line feeds.
func unifyLineFeed(s string) string {
	return strings.Replace(strings.Replace(s, "\r\n", "\n", -1), "\r", "\n", -1)
}
