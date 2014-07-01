package gohtml

import "bytes"

// A tagElement represents a tag element of an HTML document.
type tagElement struct {
	tagName     string
	startTagRaw string
	endTagRaw   string
	children    []element
  inlineParentX bool
  inlinePreviousX bool
}

// write writes a tag to the buffer.
func (e *tagElement) write(bf *bytes.Buffer, indent int) {
  if(e.hasInlineParent()) {
    writeRaw(bf, e.startTagRaw)
  } else {
    writeLine(bf, e.startTagRaw, indent)
  }
  firstChild := true
  inlinePrevious := false
	for _, c := range e.children {
    c.inlinePrevious(inlinePrevious);
		var childIndent int
      if e.inlineChildren() || (firstChild && e.goesInline()) {
        c.inlineParent(true);
      }
      firstChild = false;
		if e.endTagRaw != "" && !e.goesInline() && !e.inlineChildren() {
			childIndent = indent + 1
		} else {
			childIndent = indent
		}
    if e.goesInline() || c.hasInlineParent() || e.inlineChildren() {
      c.write(bf, childIndent)
    } else {
      c.write(bf, childIndent)
    }
    if c.goesInline() {
      inlinePrevious = true
    }
	}
	if e.endTagRaw != "" {
    if e.goesInline() || e.hasInlineParent() || e.inlineChildren() {
      writeRaw(bf, e.endTagRaw)
        if !e.goesInline() && !e.inlineChildren() {
          writeLineFeed(bf)
          writeIndent(bf, indent)
        }
    } else {
      writeLine(bf, e.endTagRaw, indent)
    }
	}
  if e.tagName == "br" {
    writeLineFeed(bf)
    writeIndent(bf, indent)
  }
}

func (e *tagElement) goesInline() bool {
  return (e.tagName == "span" ||
     e.tagName == "img" ||
     e.tagName == "br"  ||
     e.tagName == "a");
}

func (e *tagElement) inlineChildren() bool {
  return (e.tagName == "span" ||
     e.tagName == "h1"  ||
     e.tagName == "h2"  ||
     e.tagName == "h3"  ||
     e.tagName == "h4"  ||
     e.tagName == "img" ||
     e.tagName == "a");
}

// appendChild append an element to the element's children.
func (e *tagElement) appendChild(child element) {
	e.children = append(e.children, child)
}

func (e *tagElement) inlineParent(has bool) {
  e.inlineParentX = has
}

func (e *tagElement) hasInlineParent() bool {
  return e.inlineParentX
}

func (e *tagElement) inlinePrevious(has bool) {
  e.inlinePreviousX = has
}

func (e *tagElement) hasInlinePrevious() bool {
  return e.inlinePreviousX
}

