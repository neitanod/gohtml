package gohtml

import (
	"bytes"
	"strings"
)

// A textElement represents a text element of an HTML document.
type textElement struct {
	text string
	inlineParentX bool
  inlinePreviousX bool
}

// write writes a text to the buffer.
func (e *textElement) write(bf *bytes.Buffer, indent int) {
	lines := strings.Split(strings.Trim(unifyLineFeed(e.text), "\n"), "\n")
	for _, line := range lines {
    if !e.hasInlineParent() && !e.hasInlinePrevious() {
      writeLineFeed(bf)
      writeIndent(bf, indent)
    }
		bf.WriteString(line)
	}
}

func (e *textElement) writeRaw(bf *bytes.Buffer, indent int) {
	lines := strings.Split(strings.Trim(unifyLineFeed(e.text), "\n"), "\n")
	for _, line := range lines {
		writeLineFeed(bf)
		writeIndent(bf, indent)
		bf.WriteString(line)
	}
}

func (e *textElement) inlinePrevious(has bool) {
  e.inlinePreviousX = has
}

func (e *textElement) hasInlinePrevious() bool {
  return e.inlinePreviousX
}

func (e *textElement) inlineParent(has bool) {
  e.inlineParentX = has
}

func (e *textElement) hasInlineParent() bool {
  return e.inlineParentX
}

func (e *textElement) goesInline() bool {
  return false
}

func (e *textElement) inlineChildren() bool {
  return false
}
