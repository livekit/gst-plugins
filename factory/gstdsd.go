// source: gst-plugins-base

package factory

import (
	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Convert between different DSD grouping formats
type GstDsdConvert struct {
	*base.GstBaseTransform
}

func NewDsdConvert() (*GstDsdConvert, error) {
	e, err := gst.NewElement("dsdconvert")
	if err != nil {
		return nil, err
	}
	return &GstDsdConvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDsdConvertWithName(name string) (*GstDsdConvert, error) {
	e, err := gst.NewElementWithName("dsdconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstDsdConvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}
