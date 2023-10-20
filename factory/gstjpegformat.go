// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Remuxes JPEG images with markers and tags
type GstJifMux struct {
	*gst.Element
}

func NewJifMux() (*GstJifMux, error) {
	e, err := gst.NewElement("jifmux")
	if err != nil {
		return nil, err
	}
	return &GstJifMux{Element: e}, nil
}

func NewJifMuxWithName(name string) (*GstJifMux, error) {
	e, err := gst.NewElementWithName("jifmux", name)
	if err != nil {
		return nil, err
	}
	return &GstJifMux{Element: e}, nil
}

// Parse JPEG images into single-frame buffers
type GstJpegParse struct {
	*GstBaseParse
}

func NewJpegParse() (*GstJpegParse, error) {
	e, err := gst.NewElement("jpegparse")
	if err != nil {
		return nil, err
	}
	return &GstJpegParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewJpegParseWithName(name string) (*GstJpegParse, error) {
	e, err := gst.NewElementWithName("jpegparse", name)
	if err != nil {
		return nil, err
	}
	return &GstJpegParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}
