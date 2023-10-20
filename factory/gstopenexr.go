// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode EXR streams
type GstOpenEXRDec struct {
	*GstVideoDecoder
}

func NewOpenEXRDec() (*GstOpenEXRDec, error) {
	e, err := gst.NewElement("openexrdec")
	if err != nil {
		return nil, err
	}
	return &GstOpenEXRDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewOpenEXRDecWithName(name string) (*GstOpenEXRDec, error) {
	e, err := gst.NewElementWithName("openexrdec", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenEXRDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}
