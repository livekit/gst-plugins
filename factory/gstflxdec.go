// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// FLC/FLI/FLX video decoder
type GstFlxDec struct {
	*gst.Element
}

func NewFlxDec() (*GstFlxDec, error) {
	e, err := gst.NewElement("flxdec")
	if err != nil {
		return nil, err
	}
	return &GstFlxDec{Element: e}, nil
}

func NewFlxDecWithName(name string) (*GstFlxDec, error) {
	e, err := gst.NewElementWithName("flxdec", name)
	if err != nil {
		return nil, err
	}
	return &GstFlxDec{Element: e}, nil
}
