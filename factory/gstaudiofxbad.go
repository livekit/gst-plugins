// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Mixes left/right channels of stereo audio
type GstAudioChannelMix struct {
	*GstAudioFilter
}

func NewAudioChannelMix() (*GstAudioChannelMix, error) {
	e, err := gst.NewElement("audiochannelmix")
	if err != nil {
		return nil, err
	}
	return &GstAudioChannelMix{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAudioChannelMixWithName(name string) (*GstAudioChannelMix, error) {
	e, err := gst.NewElementWithName("audiochannelmix", name)
	if err != nil {
		return nil, err
	}
	return &GstAudioChannelMix{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Left channel to left channel gain
// Default: 1, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstAudioChannelMix) SetLeftToLeft(value float64) error {
	return e.Element.SetProperty("left-to-left", value)
}

func (e *GstAudioChannelMix) GetLeftToLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left-to-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Left channel to right channel gain
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstAudioChannelMix) SetLeftToRight(value float64) error {
	return e.Element.SetProperty("left-to-right", value)
}

func (e *GstAudioChannelMix) GetLeftToRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left-to-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Right channel to left channel gain
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstAudioChannelMix) SetRightToLeft(value float64) error {
	return e.Element.SetProperty("right-to-left", value)
}

func (e *GstAudioChannelMix) GetRightToLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right-to-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Right channel to right channel gain
// Default: 1, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstAudioChannelMix) SetRightToRight(value float64) error {
	return e.Element.SetProperty("right-to-right", value)
}

func (e *GstAudioChannelMix) GetRightToRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right-to-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
