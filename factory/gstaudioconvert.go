// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Convert audio to different formats
type GstAudioConvert struct {
	*base.GstBaseTransform
}

func NewAudioConvert() (*GstAudioConvert, error) {
	e, err := gst.NewElement("audioconvert")
	if err != nil {
		return nil, err
	}
	return &GstAudioConvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioConvertWithName(name string) (*GstAudioConvert, error) {
	e, err := gst.NewElementWithName("audioconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioConvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Selects between different dithering methods.
// Default: tpdf (2)
func (e *GstAudioConvert) SetDithering(value interface{}) error {
	return e.Element.SetProperty("dithering", value)
}

func (e *GstAudioConvert) GetDithering() (interface{}, error) {
	return e.Element.GetProperty("dithering")
}

// Threshold for the output bit depth at/below which to apply dithering.
// Default: 20, Min: 0, Max: 32
func (e *GstAudioConvert) SetDitheringThreshold(value uint) error {
	return e.Element.SetProperty("dithering-threshold", value)
}

func (e *GstAudioConvert) GetDitheringThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dithering-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Transformation matrix for input/output channels

func (e *GstAudioConvert) SetMixMatrix(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("mix-matrix", value)
}

func (e *GstAudioConvert) GetMixMatrix() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("mix-matrix")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Selects between different noise shaping methods.
// Default: none (0)
func (e *GstAudioConvert) SetNoiseShaping(value interface{}) error {
	return e.Element.SetProperty("noise-shaping", value)
}

func (e *GstAudioConvert) GetNoiseShaping() (interface{}, error) {
	return e.Element.GetProperty("noise-shaping")
}
