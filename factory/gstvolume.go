// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Set volume on audio/raw streams
type GstVolume struct {
	*GstAudioFilter
}

func NewVolume() (*GstVolume, error) {
	e, err := gst.NewElement("volume")
	if err != nil {
		return nil, err
	}
	return &GstVolume{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVolumeWithName(name string) (*GstVolume, error) {
	e, err := gst.NewElementWithName("volume", name)
	if err != nil {
		return nil, err
	}
	return &GstVolume{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// mute channel
// Default: false
func (e *GstVolume) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstVolume) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// volume factor, 1.0=100%%
// Default: 1, Min: 0, Max: 1.79769e+308
func (e *GstVolume) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstVolume) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
