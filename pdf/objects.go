// Copyright (C) 2011, Ross Light

package pdf

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// name is a PDF name object.
type name string

func (n name) String() string {
	return string(n)
}

func (n name) MarshalPDF() ([]byte, os.Error) {
	// TODO: escape characters
	return []byte("/" + n), nil
}

// stream is a blob of data.
type stream struct {
	Dictionary map[name]interface{}
	Bytes      []byte
}

const (
	streamBegin = "stream\r\n"
	streamEnd   = "\r\nendstream"
)

func (s stream) MarshalPDF() ([]byte, os.Error) {
	var b bytes.Buffer

	// TODO: Force Length key
	mdict, err := Marshal(s.Dictionary)
	if err != nil {
		return nil, err
	}
	b.Write(mdict)
	b.WriteString(streamBegin)
	b.Write(s.Bytes)
	b.WriteString(streamEnd)
	return b.Bytes(), nil
}

type indirectObject struct {
	Number     uint
	Generation uint
	Object     interface{}
}

const (
	objectBegin = "obj "
	objectEnd   = " endobj"
)

func (obj indirectObject) MarshalPDF() ([]byte, os.Error) {
	m, ok := obj.Object.(Marshaler)
	if !ok {
		return nil, fmt.Errorf("indirect object %d %d does not implement Marshaler", obj.Number, obj.Generation)
	}
	data, err := m.MarshalPDF()
	if err != nil {
		return nil, err
	}

	mn, mg := strconv.Uitoa(obj.Number), strconv.Uitoa(obj.Generation)
	result := make([]byte, 0, len(mn)+1+len(mg)+len(objectBegin)+len(data)+len(objectEnd))
	result = append(result, mn...)
	result = append(result, ' ')
	result = append(result, mg...)
	result = append(result, objectBegin...)
	result = append(result, data...)
	result = append(result, objectEnd...)
	return result, nil
}
