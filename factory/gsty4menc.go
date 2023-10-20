// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Encodes a YUV frame into the yuv4mpeg format (mjpegtools)
type GstY4mEncode struct {
	*GstVideoEncoder
}

func NewY4mEncode() (*GstY4mEncode, error) {
	e, err := gst.NewElement("y4menc")
	if err != nil {
		return nil, err
	}
	return &GstY4mEncode{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewY4mEncodeWithName(name string) (*GstY4mEncode, error) {
	e, err := gst.NewElementWithName("y4menc", name)
	if err != nil {
		return nil, err
	}
	return &GstY4mEncode{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}
