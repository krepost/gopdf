package pdf

type Font struct {
	pdfName     name
	pdfDict     Reference
	toCodePoint map[rune]byte
	glyphWidth  map[rune]int
}

// IsBuiltinEncoding returns true if encoding is supported by package pdf.
func IsBuiltinEncoding(encoding name) bool {
	if encoding == StandardEncoding {
		return true
	}
	_, ok := fontEncodings[encoding]
	return ok
}

// IsBuiltinFont returns true if font is one of the built-in PDF fonts.
func IsBuiltinFont(font name) bool {
	_, ok := fontMetrics[font]
	return ok
}

var fontEncodings = map[name]map[rune]byte{
	WinAnsiEncoding: map[rune]byte{
		' ':      0x20, // space
		'!':      0x21, // exclam
		'"':      0x22, // quotedbl
		'#':      0x23, // numbersign
		'$':      0x24, // dollar
		'%':      0x25, // percent
		'&':      0x26, // ampersand
		'\'':     0x27, // quotesingle
		'(':      0x28, // parenleft
		')':      0x29, // parenright
		'*':      0x2a, // asterisk
		'+':      0x2b, // plus
		',':      0x2c, // comma
		'-':      0x2d, // hyphen
		'.':      0x2e, // period
		'/':      0x2f, // slash
		'0':      0x30, // zero
		'1':      0x31, // one
		'2':      0x32, // two
		'3':      0x33, // three
		'4':      0x34, // four
		'5':      0x35, // five
		'6':      0x36, // six
		'7':      0x37, // seven
		'8':      0x38, // eight
		'9':      0x39, // nine
		':':      0x3a, // colon
		';':      0x3b, // semicolon
		'<':      0x3c, // less
		'=':      0x3d, // equal
		'>':      0x3e, // greater
		'?':      0x3f, // question
		'@':      0x40, // at
		'A':      0x41,
		'B':      0x42,
		'C':      0x43,
		'D':      0x44,
		'E':      0x45,
		'F':      0x46,
		'G':      0x47,
		'H':      0x48,
		'I':      0x49,
		'J':      0x4a,
		'K':      0x4b,
		'L':      0x4c,
		'M':      0x4d,
		'N':      0x4e,
		'O':      0x4f,
		'P':      0x50,
		'Q':      0x51,
		'R':      0x52,
		'S':      0x53,
		'T':      0x54,
		'U':      0x55,
		'V':      0x56,
		'W':      0x57,
		'X':      0x58,
		'Y':      0x59,
		'Z':      0x5a,
		'[':      0x5b, // bracketleft
		'\\':     0x5c, // backslash
		']':      0x5d, // bracketright
		'^':      0x5e, // asciicircum
		'_':      0x5f, // underscore
		'`':      0x60, // grave
		'a':      0x61,
		'b':      0x62,
		'c':      0x63,
		'd':      0x64,
		'e':      0x65,
		'f':      0x66,
		'g':      0x67,
		'h':      0x68,
		'i':      0x69,
		'j':      0x6a,
		'k':      0x6b,
		'l':      0x6c,
		'm':      0x6d,
		'n':      0x6e,
		'o':      0x6f,
		'p':      0x70,
		'q':      0x71,
		'r':      0x72,
		's':      0x73,
		't':      0x74,
		'u':      0x75,
		'v':      0x76,
		'w':      0x77,
		'x':      0x78,
		'y':      0x79,
		'z':      0x7a,
		'{':      0x7b, // braceleft
		'|':      0x7c, // bar
		'}':      0x7d, // braceright
		'~':      0x7e, // asciitilde
		'\u20ac': 0x80, // Euro
		'\u201a': 0x82, // quotesinglbase
		'\u0192': 0x83, // florin
		'\u201e': 0x84, // quotedblbase
		'\u2026': 0x85, // ellipsis
		'\u2020': 0x86, // dagger
		'\u2021': 0x87, // daggerdbl
		'\u02c6': 0x88, // circumflex
		'\u2030': 0x89, // perthousand
		'\u0160': 0x8a, // Scaron
		'\u2039': 0x8b, // guilsinglleft
		'\u0152': 0x8c, // OE
		'\u017d': 0x8e, // Zcaron
		'\u2018': 0x91, // quoteleft
		'\u2019': 0x92, // quoteright
		'\u201c': 0x93, // quotedblleft
		'\u201d': 0x94, // quotedblright
		'\u2022': 0x95, // bullet
		'\u2013': 0x96, // endash
		'\u2014': 0x97, // emdash
		'\u02dc': 0x98, // tilde
		'\u2122': 0x99, // trademark
		'\u0161': 0x9a, // scaron
		'\u203a': 0x9b, // guilsinglright
		'\u0153': 0x9c, // oe
		'\u017e': 0x9e, // zcaron
		'\u0178': 0x9f, // Ydieresis
		'\u00a1': 0xa1, // exclamdown
		'\u00a2': 0xa2, // cent
		'\u00a3': 0xa3, // sterling
		'\u00a4': 0xa4, // currency
		'\u00a5': 0xa5, // yen
		'\u00a6': 0xa6, // brokenbar
		'\u00a7': 0xa7, // section
		'\u00a8': 0xa8, // dieresis
		'\u00a9': 0xa9, // copyright
		'\u00aa': 0xaa, // ordfeminine
		'\u00ab': 0xab, // guillemotleft
		'\u00ac': 0xac, // logicalnot
		'\u00ae': 0xae, // registered
		'\u00af': 0xaf, // macron
		'\u00b0': 0xb0, // degree
		'\u00b1': 0xb1, // plusminus
		'\u00b2': 0xb2, // twosuperior
		'\u00b3': 0xb3, // threesuperior
		'\u00b4': 0xb4, // acute
		'\u03bc': 0xb5, // mu
		'\u00b6': 0xb6, // paragraph
		'\u00b7': 0xb7, // periodcentered
		'\u00b8': 0xb8, // cedilla
		'\u00b9': 0xb9, // onesuperior
		'\u00ba': 0xba, // ordmasculine
		'\u00bb': 0xbb, // guillemotright
		'\u00bc': 0xbc, // onequarter
		'\u00bd': 0xbd, // onehalf
		'\u00be': 0xbe, // threequarters
		'\u00bf': 0xbf, // questiondown
		'\u00c0': 0xc0, // Agrave
		'\u00c1': 0xc1, // Aacute
		'\u00c2': 0xc2, // Acircumflex
		'\u00c3': 0xc3, // Atilde
		'\u00c4': 0xc4, // Adieresis
		'\u00c5': 0xc5, // Aring
		'\u00c6': 0xc6, // AE
		'\u00c7': 0xc7, // Ccedilla
		'\u00c8': 0xc8, // Egrave
		'\u00c9': 0xc9, // Eacute
		'\u00ca': 0xca, // Ecircumflex
		'\u00cb': 0xcb, // Edieresis
		'\u00cc': 0xcc, // Igrave
		'\u00cd': 0xcd, // Iacute
		'\u00ce': 0xce, // Icircumflex
		'\u00cf': 0xcf, // Idieresis
		'\u00d0': 0xd0, // Eth
		'\u00d1': 0xd1, // Ntilde
		'\u00d2': 0xd2, // Ograve
		'\u00d3': 0xd3, // Oacute
		'\u00d4': 0xd4, // Ocircumflex
		'\u00d5': 0xd5, // Otilde
		'\u00d6': 0xd6, // Odieresis
		'\u00d7': 0xd7, // multiply
		'\u00d8': 0xd8, // Oslash
		'\u00d9': 0xd9, // Ugrave
		'\u00da': 0xda, // Uacute
		'\u00db': 0xdb, // Ucircumflex
		'\u00dc': 0xdc, // Udieresis
		'\u00dd': 0xdd, // Yacute
		'\u00de': 0xde, // Thorn
		'\u00df': 0xdf, // germandbls
		'\u00e0': 0xe0, // agrave
		'\u00e1': 0xe1, // aacute
		'\u00e2': 0xe2, // acircumflex
		'\u00e3': 0xe3, // atilde
		'\u00e4': 0xe4, // adieresis
		'\u00e5': 0xe5, // aring
		'\u00e6': 0xe6, // ae
		'\u00e7': 0xe7, // ccedilla
		'\u00e8': 0xe8, // egrave
		'\u00e9': 0xe9, // eacute
		'\u00ea': 0xea, // ecircumflex
		'\u00eb': 0xeb, // edieresis
		'\u00ec': 0xec, // igrave
		'\u00ed': 0xed, // iacute
		'\u00ee': 0xee, // icircumflex
		'\u00ef': 0xef, // idieresis
		'\u00f0': 0xf0, // eth
		'\u00f1': 0xf1, // ntilde
		'\u00f2': 0xf2, // ograve
		'\u00f3': 0xf3, // oacute
		'\u00f4': 0xf4, // ocircumflex
		'\u00f5': 0xf5, // otilde
		'\u00f6': 0xf6, // odieresis
		'\u00f7': 0xf7, // divide
		'\u00f8': 0xf8, // oslash
		'\u00f9': 0xf9, // ugrave
		'\u00fa': 0xfa, // uacute
		'\u00fb': 0xfb, // ucircumflex
		'\u00fc': 0xfc, // udieresis
		'\u00fd': 0xfd, // yacute
		'\u00fe': 0xfe, // thorn
		'\u00ff': 0xff, // ydieresis
	},

	MacRomanEncoding: map[rune]byte{
		' ':      0x20, // space
		'!':      0x21, // exclam
		'"':      0x22, // quotedbl
		'#':      0x23, // numbersign
		'$':      0x24, // dollar
		'%':      0x25, // percent
		'&':      0x26, // ampersand
		'\'':     0x27, // quotesingle
		'(':      0x28, // parenleft
		')':      0x29, // parenright
		'*':      0x2a, // asterisk
		'+':      0x2b, // plus
		',':      0x2c, // comma
		'-':      0x2d, // hyphen
		'.':      0x2e, // period
		'/':      0x2f, // slash
		'0':      0x30, // zero
		'1':      0x31, // one
		'2':      0x32, // two
		'3':      0x33, // three
		'4':      0x34, // four
		'5':      0x35, // five
		'6':      0x36, // six
		'7':      0x37, // seven
		'8':      0x38, // eight
		'9':      0x39, // nine
		':':      0x3a, // colon
		';':      0x3b, // semicolon
		'<':      0x3c, // less
		'=':      0x3d, // equal
		'>':      0x3e, // greater
		'?':      0x3f, // question
		'@':      0x40, // at
		'A':      0x41,
		'B':      0x42,
		'C':      0x43,
		'D':      0x44,
		'E':      0x45,
		'F':      0x46,
		'G':      0x47,
		'H':      0x48,
		'I':      0x49,
		'J':      0x4a,
		'K':      0x4b,
		'L':      0x4c,
		'M':      0x4d,
		'N':      0x4e,
		'O':      0x4f,
		'P':      0x50,
		'Q':      0x51,
		'R':      0x52,
		'S':      0x53,
		'T':      0x54,
		'U':      0x55,
		'V':      0x56,
		'W':      0x57,
		'X':      0x58,
		'Y':      0x59,
		'Z':      0x5a,
		'[':      0x5b, // bracketleft
		'\\':     0x5c, // backslash
		']':      0x5d, // bracketright
		'^':      0x5e, // asciicircum
		'_':      0x5f, // underscore
		'`':      0x60, // grave
		'a':      0x61,
		'b':      0x62,
		'c':      0x63,
		'd':      0x64,
		'e':      0x65,
		'f':      0x66,
		'g':      0x67,
		'h':      0x68,
		'i':      0x69,
		'j':      0x6a,
		'k':      0x6b,
		'l':      0x6c,
		'm':      0x6d,
		'n':      0x6e,
		'o':      0x6f,
		'p':      0x70,
		'q':      0x71,
		'r':      0x72,
		's':      0x73,
		't':      0x74,
		'u':      0x75,
		'v':      0x76,
		'w':      0x77,
		'x':      0x78,
		'y':      0x79,
		'z':      0x7a,
		'{':      0x7b, // braceleft
		'|':      0x7c, // bar
		'}':      0x7d, // braceright
		'~':      0x7e, // asciitilde
		'\u00c4': 0x80, // Adieresis
		'\u00c5': 0x81, // Aring
		'\u00c7': 0x82, // Ccedilla
		'\u00c9': 0x83, // Eacute
		'\u00d1': 0x84, // Ntilde
		'\u00d6': 0x85, // Odieresis
		'\u00dc': 0x86, // Udieresis
		'\u00e1': 0x87, // aacute
		'\u00e0': 0x88, // agrave
		'\u00e2': 0x89, // acircumflex
		'\u00e4': 0x8a, // adieresis
		'\u00e3': 0x8b, // atilde
		'\u00e5': 0x8c, // aring
		'\u00e7': 0x8d, // ccedilla
		'\u00e9': 0x8e, // eacute
		'\u00e8': 0x8f, // egrave
		'\u00ea': 0x90, // ecircumflex
		'\u00eb': 0x91, // edieresis
		'\u00ed': 0x92, // iacute
		'\u00ec': 0x93, // igrave
		'\u00ee': 0x94, // icircumflex
		'\u00ef': 0x95, // idieresis
		'\u00f1': 0x96, // ntilde
		'\u00f3': 0x97, // oacute
		'\u00f2': 0x98, // ograve
		'\u00f4': 0x99, // ocircumflex
		'\u00f6': 0x9a, // odieresis
		'\u00f5': 0x9b, // otilde
		'\u00fa': 0x9c, // uacute
		'\u00f9': 0x9d, // ugrave
		'\u00fb': 0x9e, // ucircumflex
		'\u00fc': 0x9f, // udieresis
		'\u2020': 0xa0, // dagger
		'\u00b0': 0xa1, // degree
		'\u00a2': 0xa2, // cent
		'\u00a3': 0xa3, // sterling
		'\u00a7': 0xa4, // section
		'\u2022': 0xa5, // bullet
		'\u00b6': 0xa6, // paragraph
		'\u00df': 0xa7, // germandbls
		'\u00ae': 0xa8, // registered
		'\u00a9': 0xa9, // copyright
		'\u2122': 0xaa, // trademark
		'\u00b4': 0xab, // acute
		'\u00a8': 0xac, // dieresis
		'\u00c6': 0xae, // AE
		'\u00d8': 0xaf, // Oslash
		'\u00b1': 0xb1, // plusminus
		'\u00a5': 0xb4, // yen
		'\u03bc': 0xb5, // mu
		'\u00aa': 0xbb, // ordfeminine
		'\u00ba': 0xbc, // ordmasculine
		'\u00e6': 0xbe, // ae
		'\u00f8': 0xbf, // oslash
		'\u00bf': 0xc0, // questiondown
		'\u00a1': 0xc1, // exclamdown
		'\u00ac': 0xc2, // logicalnot
		'\u0192': 0xc4, // florin
		'\u00ab': 0xc7, // guillemotleft
		'\u00bb': 0xc8, // guillemotright
		'\u2026': 0xc9, // ellipsis
		'\u00c0': 0xcb, // Agrave
		'\u00c3': 0xcc, // Atilde
		'\u00d5': 0xcd, // Otilde
		'\u0152': 0xce, // OE
		'\u0153': 0xcf, // oe
		'\u2013': 0xd0, // endash
		'\u2014': 0xd1, // emdash
		'\u201c': 0xd2, // quotedblleft
		'\u201d': 0xd3, // quotedblright
		'\u2018': 0xd4, // quoteleft
		'\u2019': 0xd5, // quoteright
		'\u00f7': 0xd6, // divide
		'\u00ff': 0xd8, // ydieresis
		'\u0178': 0xd9, // Ydieresis
		'\u2044': 0xda, // fraction
		'\u00a4': 0xdb, // currency
		'\u2039': 0xdc, // guilsinglleft
		'\u203a': 0xdd, // guilsinglright
		'\ufb01': 0xde, // fi
		'\ufb02': 0xdf, // fl
		'\u2021': 0xe0, // daggerdbl
		'\u00b7': 0xe1, // periodcentered
		'\u201a': 0xe2, // quotesinglbase
		'\u201e': 0xe3, // quotedblbase
		'\u2030': 0xe4, // perthousand
		'\u00c2': 0xe5, // Acircumflex
		'\u00ca': 0xe6, // Ecircumflex
		'\u00c1': 0xe7, // Aacute
		'\u00cb': 0xe8, // Edieresis
		'\u00c8': 0xe9, // Egrave
		'\u00cd': 0xea, // Iacute
		'\u00ce': 0xeb, // Icircumflex
		'\u00cf': 0xec, // Idieresis
		'\u00cc': 0xed, // Igrave
		'\u00d3': 0xee, // Oacute
		'\u00d4': 0xef, // Ocircumflex
		'\u00d2': 0xf1, // Ograve
		'\u00da': 0xf2, // Uacute
		'\u00db': 0xf3, // Ucircumflex
		'\u00d9': 0xf4, // Ugrave
		'\u0131': 0xf5, // dotlessi
		'\u02c6': 0xf6, // circumflex
		'\u02dc': 0xf7, // tilde
		'\u00af': 0xf8, // macron
		'\u02d8': 0xf9, // breve
		'\u02d9': 0xfa, // dotaccent
		'\u02da': 0xfb, // ring
		'\u00b8': 0xfc, // cedilla
		'\u02dd': 0xfd, // hungarumlaut
		'\u02db': 0xfe, // ogonek
		'\u02c7': 0xff, // caron
	},

	PDFDocEncoding: map[rune]byte{
		'\u02d8': 0x18, // breve
		'\u02c7': 0x19, // caron
		'\u02c6': 0x1a, // circumflex
		'\u02d9': 0x1b, // dotaccent
		'\u02dd': 0x1c, // hungarumlaut
		'\u02db': 0x1d, // ogonek
		'\u02da': 0x1e, // ring
		'\u02dc': 0x1f, // tilde
		' ':      0x20, // space
		'!':      0x21, // exclam
		'"':      0x22, // quotedbl
		'#':      0x23, // numbersign
		'$':      0x24, // dollar
		'%':      0x25, // percent
		'&':      0x26, // ampersand
		'\'':     0x27, // quotesingle
		'(':      0x28, // parenleft
		')':      0x29, // parenright
		'*':      0x2a, // asterisk
		'+':      0x2b, // plus
		',':      0x2c, // comma
		'-':      0x2d, // hyphen
		'.':      0x2e, // period
		'/':      0x2f, // slash
		'0':      0x30, // zero
		'1':      0x31, // one
		'2':      0x32, // two
		'3':      0x33, // three
		'4':      0x34, // four
		'5':      0x35, // five
		'6':      0x36, // six
		'7':      0x37, // seven
		'8':      0x38, // eight
		'9':      0x39, // nine
		':':      0x3a, // colon
		';':      0x3b, // semicolon
		'<':      0x3c, // less
		'=':      0x3d, // equal
		'>':      0x3e, // greater
		'?':      0x3f, // question
		'@':      0x40, // at
		'A':      0x41,
		'B':      0x42,
		'C':      0x43,
		'D':      0x44,
		'E':      0x45,
		'F':      0x46,
		'G':      0x47,
		'H':      0x48,
		'I':      0x49,
		'J':      0x4a,
		'K':      0x4b,
		'L':      0x4c,
		'M':      0x4d,
		'N':      0x4e,
		'O':      0x4f,
		'P':      0x50,
		'Q':      0x51,
		'R':      0x52,
		'S':      0x53,
		'T':      0x54,
		'U':      0x55,
		'V':      0x56,
		'W':      0x57,
		'X':      0x58,
		'Y':      0x59,
		'Z':      0x5a,
		'[':      0x5b, // bracketleft
		'\\':     0x5c, // backslash
		']':      0x5d, // bracketright
		'^':      0x5e, // asciicircum
		'_':      0x5f, // underscore
		'`':      0x60, // grave
		'a':      0x61,
		'b':      0x62,
		'c':      0x63,
		'd':      0x64,
		'e':      0x65,
		'f':      0x66,
		'g':      0x67,
		'h':      0x68,
		'i':      0x69,
		'j':      0x6a,
		'k':      0x6b,
		'l':      0x6c,
		'm':      0x6d,
		'n':      0x6e,
		'o':      0x6f,
		'p':      0x70,
		'q':      0x71,
		'r':      0x72,
		's':      0x73,
		't':      0x74,
		'u':      0x75,
		'v':      0x76,
		'w':      0x77,
		'x':      0x78,
		'y':      0x79,
		'z':      0x7a,
		'{':      0x7b, // braceleft
		'|':      0x7c, // bar
		'}':      0x7d, // braceright
		'~':      0x7e, // asciitilde
		'\u2022': 0x80, // bullet
		'\u2020': 0x81, // dagger
		'\u2021': 0x82, // daggerdbl
		'\u2026': 0x83, // ellipsis
		'\u2014': 0x84, // emdash
		'\u2013': 0x85, // endash
		'\u0192': 0x86, // florin
		'\u2044': 0x87, // fraction
		'\u2039': 0x88, // guilsinglleft
		'\u203a': 0x89, // guilsinglright
		'\u2212': 0x8a, // minus
		'\u2030': 0x8b, // perthousand
		'\u201e': 0x8c, // quotedblbase
		'\u201c': 0x8d, // quotedblleft
		'\u201d': 0x8e, // quotedblright
		'\u2018': 0x8f, // quoteleft
		'\u2019': 0x90, // quoteright
		'\u201a': 0x91, // quotesinglbase
		'\u2122': 0x92, // trademark
		'\ufb01': 0x93, // fi
		'\ufb02': 0x94, // fl
		'\u0141': 0x95, // Lslash
		'\u0152': 0x96, // OE
		'\u0160': 0x97, // Scaron
		'\u0178': 0x98, // Ydieresis
		'\u017d': 0x99, // Zcaron
		'\u0131': 0x9a, // dotlessi
		'\u0142': 0x9b, // lslash
		'\u0153': 0x9c, // oe
		'\u0161': 0x9d, // scaron
		'\u017e': 0x9e, // zcaron
		'\u20ac': 0xa0, // Euro
		'\u00a1': 0xa1, // exclamdown
		'\u00a2': 0xa2, // cent
		'\u00a3': 0xa3, // sterling
		'\u00a4': 0xa4, // currency
		'\u00a5': 0xa5, // yen
		'\u00a6': 0xa6, // brokenbar
		'\u00a7': 0xa7, // section
		'\u00a8': 0xa8, // dieresis
		'\u00a9': 0xa9, // copyright
		'\u00aa': 0xaa, // ordfeminine
		'\u00ab': 0xab, // guillemotleft
		'\u00ac': 0xac, // logicalnot
		'\u00ae': 0xae, // registered
		'\u00af': 0xaf, // macron
		'\u00b0': 0xb0, // degree
		'\u00b1': 0xb1, // plusminus
		'\u00b2': 0xb2, // twosuperior
		'\u00b3': 0xb3, // threesuperior
		'\u00b4': 0xb4, // acute
		'\u03bc': 0xb5, // mu
		'\u00b6': 0xb6, // paragraph
		'\u00b7': 0xb7, // periodcentered
		'\u00b8': 0xb8, // cedilla
		'\u00b9': 0xb9, // onesuperior
		'\u00ba': 0xba, // ordmasculine
		'\u00bb': 0xbb, // guillemotright
		'\u00bc': 0xbc, // onequarter
		'\u00bd': 0xbd, // onehalf
		'\u00be': 0xbe, // threequarters
		'\u00bf': 0xbf, // questiondown
		'\u00c0': 0xc0, // Agrave
		'\u00c1': 0xc1, // Aacute
		'\u00c2': 0xc2, // Acircumflex
		'\u00c3': 0xc3, // Atilde
		'\u00c4': 0xc4, // Adieresis
		'\u00c5': 0xc5, // Aring
		'\u00c6': 0xc6, // AE
		'\u00c7': 0xc7, // Ccedilla
		'\u00c8': 0xc8, // Egrave
		'\u00c9': 0xc9, // Eacute
		'\u00ca': 0xca, // Ecircumflex
		'\u00cb': 0xcb, // Edieresis
		'\u00cc': 0xcc, // Igrave
		'\u00cd': 0xcd, // Iacute
		'\u00ce': 0xce, // Icircumflex
		'\u00cf': 0xcf, // Idieresis
		'\u00d0': 0xd0, // Eth
		'\u00d1': 0xd1, // Ntilde
		'\u00d2': 0xd2, // Ograve
		'\u00d3': 0xd3, // Oacute
		'\u00d4': 0xd4, // Ocircumflex
		'\u00d5': 0xd5, // Otilde
		'\u00d6': 0xd6, // Odieresis
		'\u00d7': 0xd7, // multiply
		'\u00d8': 0xd8, // Oslash
		'\u00d9': 0xd9, // Ugrave
		'\u00da': 0xda, // Uacute
		'\u00db': 0xdb, // Ucircumflex
		'\u00dc': 0xdc, // Udieresis
		'\u00dd': 0xdd, // Yacute
		'\u00de': 0xde, // Thorn
		'\u00df': 0xdf, // germandbls
		'\u00e0': 0xe0, // agrave
		'\u00e1': 0xe1, // aacute
		'\u00e2': 0xe2, // acircumflex
		'\u00e3': 0xe3, // atilde
		'\u00e4': 0xe4, // adieresis
		'\u00e5': 0xe5, // aring
		'\u00e6': 0xe6, // ae
		'\u00e7': 0xe7, // ccedilla
		'\u00e8': 0xe8, // egrave
		'\u00e9': 0xe9, // eacute
		'\u00ea': 0xea, // ecircumflex
		'\u00eb': 0xeb, // edieresis
		'\u00ec': 0xec, // igrave
		'\u00ed': 0xed, // iacute
		'\u00ee': 0xee, // icircumflex
		'\u00ef': 0xef, // idieresis
		'\u00f0': 0xf0, // eth
		'\u00f1': 0xf1, // ntilde
		'\u00f2': 0xf2, // ograve
		'\u00f3': 0xf3, // oacute
		'\u00f4': 0xf4, // ocircumflex
		'\u00f5': 0xf5, // otilde
		'\u00f6': 0xf6, // odieresis
		'\u00f7': 0xf7, // divide
		'\u00f8': 0xf8, // oslash
		'\u00f9': 0xf9, // ugrave
		'\u00fa': 0xfa, // uacute
		'\u00fb': 0xfb, // ucircumflex
		'\u00fc': 0xfc, // udieresis
		'\u00fd': 0xfd, // yacute
		'\u00fe': 0xfe, // thorn
		'\u00ff': 0xff, // ydieresis
	},
}
