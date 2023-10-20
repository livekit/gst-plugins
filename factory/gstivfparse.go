// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Demuxes a IVF stream
type GstIvfParse struct {
	*GstBaseParse
}

func NewIvfParse() (*GstIvfParse, error) {
	e, err := gst.NewElement("ivfparse")
	if err != nil {
		return nil, err
	}
	return &GstIvfParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewIvfParseWithName(name string) (*GstIvfParse, error) {
	e, err := gst.NewElementWithName("ivfparse", name)
	if err != nil {
		return nil, err
	}
	return &GstIvfParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}
