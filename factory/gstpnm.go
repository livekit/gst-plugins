// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes images in portable pixmap/graymap/bitmap/anymamp (PNM) format
type GstPnmdec struct {
	*GstVideoDecoder
}

func NewPnmdec() (*GstPnmdec, error) {
	e, err := gst.NewElement("pnmdec")
	if err != nil {
		return nil, err
	}
	return &GstPnmdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewPnmdecWithName(name string) (*GstPnmdec, error) {
	e, err := gst.NewElementWithName("pnmdec", name)
	if err != nil {
		return nil, err
	}
	return &GstPnmdec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Encodes images into portable pixmap or graymap (PNM) format
type GstPnmenc struct {
	*GstVideoEncoder
}

func NewPnmenc() (*GstPnmenc, error) {
	e, err := gst.NewElement("pnmenc")
	if err != nil {
		return nil, err
	}
	return &GstPnmenc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewPnmencWithName(name string) (*GstPnmenc, error) {
	e, err := gst.NewElementWithName("pnmenc", name)
	if err != nil {
		return nil, err
	}
	return &GstPnmenc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// The output will be ASCII encoded
// Default: false
func (e *GstPnmenc) SetAscii(value bool) error {
	return e.Element.SetProperty("ascii", value)
}

func (e *GstPnmenc) GetAscii() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ascii")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
