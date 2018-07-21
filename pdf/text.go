// Copyright (C) 2011, Ross Light

package pdf

import (
	"bytes"
)

// Text is a PDF text object.  The zero value is an empty text object.
type Text struct {
	buf   bytes.Buffer
	fonts map[name]*Font

	x, y        Unit
	currFont    *Font
	currSize    Unit
	currLeading Unit
}

// Width computes the width of a string in the given font and font size.
func (font *Font) Width(s string, fontSize Unit) Unit {
	width := Unit(0)
	for _, r := range s {
		if w, ok := font.glyphWidth[r]; ok {
			width += Unit(w)
		}
	}
	return width * fontSize / 1000
}

// CodePoints encodes a string as a sequence of code points in the given font.
func (font *Font) CodePoints(s string) []byte {
	encoded := []byte{}
	for _, r := range s {
		if b, ok := font.toCodePoint[r]; ok {
			encoded = append(encoded, b)
		}
	}
	return encoded
}

// Text adds a string to the text object.
func (text *Text) Text(s string) {
	text.x += text.currFont.Width(s, text.currSize)
	encoded := text.currFont.CodePoints(s)
	writeCommand(&text.buf, "Tj", string(encoded))
}

// UseFont changes the current font with given size and leading.
func (text *Text) UseFont(font *Font, size Unit, leading Unit) {
	if text.fonts == nil {
		text.fonts = make(map[name]*Font)
	}
	if f, ok := text.fonts[font.pdfName]; ok {
		font = f
	} else {
		text.fonts[font.pdfName] = font
	}
	text.currFont = font
	text.currSize = size
	writeCommand(&text.buf, "Tf", font.pdfName, size)
	text.SetLeading(leading)
}

const defaultLeadingScalar = 1.2

// SetFont changes the current font to a standard font.  This also changes the
// leading to 1.2 times the font size.
func (text *Text) SetFont(fontName string, size Unit) {
	// font.pdfName==fontName for fonts in default encoding. We don't set
	// pdfDict here but instead create it lazily when the text is added to
	// the canvas. At that time, we have access to the Document object,
	// which is needed in order to store the document in the correct dictionary.
	f := Font{
		pdfName:     name(fontName),
		toCodePoint: map[rune]byte{},
		glyphWidth:  map[rune]int{},
	}
	// toCodePoint is the identity function for default encoding.
	// glyphWidth contains widths for all encoded glyphs.
	for _, m := range fontMetrics[name(fontName)] {
		if m.codePoint != -1 {
			f.toCodePoint[m.rune] = byte(m.codePoint)
			f.glyphWidth[m.rune] = m.width
		}
	}
	text.UseFont(&f, size, size*defaultLeadingScalar)
}

// SetLeading changes the amount of space between lines.
func (text *Text) SetLeading(leading Unit) {
	writeCommand(&text.buf, "TL", leading)
	text.currLeading = leading
}

// NextLine advances the current text position to the next line, based on the
// current leading.
func (text *Text) NextLine() {
	writeCommand(&text.buf, "T*")
	text.x = 0
	text.y -= text.currLeading
}

// NextLineOffset moves the current text position to an offset relative to the
// beginning of the line.
func (text *Text) NextLineOffset(tx, ty Unit) {
	writeCommand(&text.buf, "Td", tx, ty)
	text.x = tx
	text.y += ty
}

// X returns the current x position of the text cursor.
func (text *Text) X() Unit {
	return text.x
}

// Y returns the current y position of the text cursor.
func (text *Text) Y() Unit {
	return text.y
}

// Standard 14 fonts
const (
	Courier            = "Courier"
	CourierBold        = "Courier-Bold"
	CourierOblique     = "Courier-Oblique"
	CourierBoldOblique = "Courier-BoldOblique"

	Helvetica            = "Helvetica"
	HelveticaBold        = "Helvetica-Bold"
	HelveticaOblique     = "Helvetica-Oblique"
	HelveticaBoldOblique = "Helvetica-BoldOblique"

	Symbol = "Symbol"

	Times           = "Times-Roman"
	TimesBold       = "Times-Bold"
	TimesItalic     = "Times-Italic"
	TimesBoldItalic = "Times-BoldItalic"

	ZapfDingbats = "ZapfDingbats"
)

// Supported font encodings
const (
	StandardEncoding = ""
	WinAnsiEncoding  = "WinAnsiEncoding"
	MacRomanEncoding = "MacRomanEncoding"
	PDFDocEncoding   = "PDFDocEncoding"
)
