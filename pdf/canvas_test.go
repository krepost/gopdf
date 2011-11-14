// Copyright (C) 2011, Ross Light

package pdf

import (
	"testing"
)

const textExpectedOutput = `/Helvetica 12 Tf
14 TL
(Hello, World!) Tj
T*
(This is SPARTA!!1!) Tj
`

func TestText(t *testing.T) {
	text := new(Text)
	text.SetFont(Helvetica, 12)
	text.SetLeading(14)
	text.Show("Hello, World!")
	text.NextLine()
	text.Show("This is SPARTA!!1!")

	if text.buf.String() != textExpectedOutput {
		t.Errorf("Output was %q, expected %q", text.buf.String(), textExpectedOutput)
	}

	if len(text.fonts) == 1 {
		if !text.fonts[Helvetica] {
			t.Error("Helvetica missing from fonts")
		}
	} else {
		t.Errorf("Got %d fonts, expected %d", len(text.fonts), 1)
	}
}