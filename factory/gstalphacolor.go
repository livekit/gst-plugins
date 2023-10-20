// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// ARGB from/to AYUV colorspace conversion preserving the alpha channel
type GstAlphaColor struct {
	*GstVideoFilter
}

func NewAlphaColor() (*GstAlphaColor, error) {
	e, err := gst.NewElement("alphacolor")
	if err != nil {
		return nil, err
	}
	return &GstAlphaColor{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAlphaColorWithName(name string) (*GstAlphaColor, error) {
	e, err := gst.NewElementWithName("alphacolor", name)
	if err != nil {
		return nil, err
	}
	return &GstAlphaColor{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}
