// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Converts video/x-bayer to video/x-raw
type GstBayer2RGB struct {
	*base.GstBaseTransform
}

func NewBayer2RGB() (*GstBayer2RGB, error) {
	e, err := gst.NewElement("bayer2rgb")
	if err != nil {
		return nil, err
	}
	return &GstBayer2RGB{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBayer2RGBWithName(name string) (*GstBayer2RGB, error) {
	e, err := gst.NewElementWithName("bayer2rgb", name)
	if err != nil {
		return nil, err
	}
	return &GstBayer2RGB{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Converts video/x-raw to video/x-bayer
type GstRGB2Bayer struct {
	*base.GstBaseTransform
}

func NewRGB2Bayer() (*GstRGB2Bayer, error) {
	e, err := gst.NewElement("rgb2bayer")
	if err != nil {
		return nil, err
	}
	return &GstRGB2Bayer{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRGB2BayerWithName(name string) (*GstRGB2Bayer, error) {
	e, err := gst.NewElementWithName("rgb2bayer", name)
	if err != nil {
		return nil, err
	}
	return &GstRGB2Bayer{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}
