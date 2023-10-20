// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes compressed streams
type GstBz2dec struct {
	*gst.Element
}

func NewBz2dec() (*GstBz2dec, error) {
	e, err := gst.NewElement("bz2dec")
	if err != nil {
		return nil, err
	}
	return &GstBz2dec{Element: e}, nil
}

func NewBz2decWithName(name string) (*GstBz2dec, error) {
	e, err := gst.NewElementWithName("bz2dec", name)
	if err != nil {
		return nil, err
	}
	return &GstBz2dec{Element: e}, nil
}

// Buffer size
// Default: 1024, Min: 1, Max: -1
func (e *GstBz2dec) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstBz2dec) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Size of first buffer (used to determine the mime type of the uncompressed data)
// Default: 1024, Min: 1, Max: -1
func (e *GstBz2dec) SetFirstBufferSize(value uint) error {
	return e.Element.SetProperty("first-buffer-size", value)
}

func (e *GstBz2dec) GetFirstBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("first-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Compresses streams
type GstBz2enc struct {
	*gst.Element
}

func NewBz2enc() (*GstBz2enc, error) {
	e, err := gst.NewElement("bz2enc")
	if err != nil {
		return nil, err
	}
	return &GstBz2enc{Element: e}, nil
}

func NewBz2encWithName(name string) (*GstBz2enc, error) {
	e, err := gst.NewElementWithName("bz2enc", name)
	if err != nil {
		return nil, err
	}
	return &GstBz2enc{Element: e}, nil
}

// Buffer size
// Default: 1024, Min: 1, Max: -1
func (e *GstBz2enc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstBz2enc) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Block size
// Default: 6, Min: 1, Max: 9
func (e *GstBz2enc) SetBlockSize(value uint) error {
	return e.Element.SetProperty("block-size", value)
}

func (e *GstBz2enc) GetBlockSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("block-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
