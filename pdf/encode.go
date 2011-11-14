// Copyright (C) 2011, Ross Light

package pdf

import (
	"fmt"
	"io"
	"os"
)

// An Encoder writes the PDF file format.
type Encoder struct {
	objects []interface{}
	root    Reference
}

type trailer struct {
	Size int
	Root Reference
}

func (enc *Encoder) Add(v interface{}) Reference {
	enc.objects = append(enc.objects, v)
	return Reference{uint(len(enc.objects)), 0}
}

const (
	header  = "%PDF-1.7" + newline + "%\x93\x8c\x8b\x9e" + newline
	newline = "\r\n"
)

const (
	crossReferenceSectionHeader    = "xref" + newline
	crossReferenceSubsectionFormat = "%d %d" + newline
	crossReferenceFormat           = "%010d %05d n" + newline
	crossReferenceFreeFormat       = "%010d %05d f" + newline
)

const trailerHeader = "trailer" + newline

const startxrefFormat = "startxref" + newline + "%d" + newline

const eofString = "%%EOF" + newline

func (enc *Encoder) Encode(wr io.Writer) os.Error {
	w := &offsetWriter{Writer: wr}

	// Write header
	if _, err := io.WriteString(w, header); err != nil {
		return err
	}

	// Write body
	objectOffsets := make([]int64, len(enc.objects))
	for i, obj := range enc.objects {
		objectOffsets[i] = w.offset
		data, err := Marshal(indirectObject{uint(i + 1), 0, obj})
		if err != nil {
			return err
		}
		if _, err = w.Write(data); err != nil {
			return err
		}
		if _, err = io.WriteString(w, newline); err != nil {
			return err
		}
	}

	// Write cross-reference table
	tableOffset := w.offset
	if _, err := io.WriteString(w, crossReferenceSectionHeader); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, crossReferenceSubsectionFormat, 0, len(enc.objects)+1); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, crossReferenceFreeFormat, len(enc.objects)+1, 65535); err != nil {
		return err
	}
	for _, offset := range objectOffsets {
		if _, err := fmt.Fprintf(w, crossReferenceFormat, offset, 0); err != nil {
			return err
		}
	}

	// Write trailer
	if _, err := io.WriteString(w, trailerHeader); err != nil {
		return err
	}
	trailerDict := trailer{
		Size: len(enc.objects) + 1,
		Root: enc.root,
	}
	trailerData, err := Marshal(trailerDict)
	if err != nil {
		return err
	}
	if _, err := w.Write(trailerData); err != nil {
		return err
	}
	if _, err := io.WriteString(w, newline); err != nil {
		return err
	}

	// Write startxref
	if fmt.Fprintf(w, startxrefFormat, tableOffset); err != nil {
		return err
	}

	// Finish file
	if _, err := io.WriteString(w, eofString); err != nil {
		return err
	}
	return nil
}

type offsetWriter struct {
	io.Writer
	offset int64
}

func (w *offsetWriter) Write(p []byte) (n int, err os.Error) {
	n, err = w.Writer.Write(p)
	w.offset += int64(n)
	return
}
