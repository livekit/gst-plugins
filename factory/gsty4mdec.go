// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Demuxes/decodes YUV4MPEG streams
type GstY4mDec struct {
	*gst.Element
}

func NewY4mDec() (*GstY4mDec, error) {
	e, err := gst.NewElement("y4mdec")
	if err != nil {
		return nil, err
	}
	return &GstY4mDec{Element: e}, nil
}

func NewY4mDecWithName(name string) (*GstY4mDec, error) {
	e, err := gst.NewElementWithName("y4mdec", name)
	if err != nil {
		return nil, err
	}
	return &GstY4mDec{Element: e}, nil
}
