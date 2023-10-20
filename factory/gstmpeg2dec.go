// source: gst-plugins-ugly

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Uses libmpeg2 to decode MPEG video streams
type GstMpeg2dec struct {
	*GstVideoDecoder
}

func NewMpeg2dec() (*GstMpeg2dec, error) {
	e, err := gst.NewElement("mpeg2dec")
	if err != nil {
		return nil, err
	}
	return &GstMpeg2dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewMpeg2decWithName(name string) (*GstMpeg2dec, error) {
	e, err := gst.NewElementWithName("mpeg2dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMpeg2dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}
