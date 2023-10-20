// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// parses opus audio streams
type GstOpusParse struct {
	*GstBaseParse
}

func NewOpusParse() (*GstOpusParse, error) {
	e, err := gst.NewElement("opusparse")
	if err != nil {
		return nil, err
	}
	return &GstOpusParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewOpusParseWithName(name string) (*GstOpusParse, error) {
	e, err := gst.NewElementWithName("opusparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOpusParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}
