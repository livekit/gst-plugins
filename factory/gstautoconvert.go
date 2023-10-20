// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Selects the right transform element based on the caps
type GstAutoConvert struct {
	*gst.Bin
}

func NewAutoConvert() (*GstAutoConvert, error) {
	e, err := gst.NewElement("autoconvert")
	if err != nil {
		return nil, err
	}
	return &GstAutoConvert{Bin: &gst.Bin{Element: e}}, nil
}

func NewAutoConvertWithName(name string) (*GstAutoConvert, error) {
	e, err := gst.NewElementWithName("autoconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoConvert{Bin: &gst.Bin{Element: e}}, nil
}

// GList of GstElementFactory objects to pick from (the element takes ownership of the list (NULL means it will go through all possible elements), can only be set once

func (e *GstAutoConvert) SetFactories(value interface{}) error {
	return e.Element.SetProperty("factories", value)
}

func (e *GstAutoConvert) GetFactories() (interface{}, error) {
	return e.Element.GetProperty("factories")
}

// Selects the right color space converter based on the caps
type GstAutoVideoConvert struct {
	*gst.Bin
}

func NewAutoVideoConvert() (*GstAutoVideoConvert, error) {
	e, err := gst.NewElement("autovideoconvert")
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoConvert{Bin: &gst.Bin{Element: e}}, nil
}

func NewAutoVideoConvertWithName(name string) (*GstAutoVideoConvert, error) {
	e, err := gst.NewElementWithName("autovideoconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstAutoVideoConvert{Bin: &gst.Bin{Element: e}}, nil
}
