// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Encode a video frame to a .png image
type GstPngEnc struct {
	*GstVideoEncoder
}

func NewPngEnc() (*GstPngEnc, error) {
	e, err := gst.NewElement("pngenc")
	if err != nil {
		return nil, err
	}
	return &GstPngEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewPngEncWithName(name string) (*GstPngEnc, error) {
	e, err := gst.NewElementWithName("pngenc", name)
	if err != nil {
		return nil, err
	}
	return &GstPngEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// PNG compression level
// Default: 6, Min: 0, Max: 9
func (e *GstPngEnc) SetCompressionLevel(value uint) error {
	return e.Element.SetProperty("compression-level", value)
}

func (e *GstPngEnc) GetCompressionLevel() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("compression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Send EOS after encoding a frame, useful for snapshots
// Default: false
func (e *GstPngEnc) SetSnapshot(value bool) error {
	return e.Element.SetProperty("snapshot", value)
}

func (e *GstPngEnc) GetSnapshot() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("snapshot")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Decode a png video frame to a raw image
type GstPngDec struct {
	*GstVideoDecoder
}

func NewPngDec() (*GstPngDec, error) {
	e, err := gst.NewElement("pngdec")
	if err != nil {
		return nil, err
	}
	return &GstPngDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewPngDecWithName(name string) (*GstPngDec, error) {
	e, err := gst.NewElementWithName("pngdec", name)
	if err != nil {
		return nil, err
	}
	return &GstPngDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}
