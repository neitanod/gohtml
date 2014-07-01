package gohtml

import "bytes"

// An element represents an HTML element.
type element interface {
	write(*bytes.Buffer, int)
  inlinePrevious(bool)
  hasInlinePrevious() bool
  inlineParent(bool)
  hasInlineParent() bool
  goesInline() bool
  inlineChildren() bool
}
