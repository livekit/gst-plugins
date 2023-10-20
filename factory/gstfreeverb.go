// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Add reverberation to audio streams
type GstFreeverb struct {
	*base.GstBaseTransform
}

func NewFreeverb() (*GstFreeverb, error) {
	e, err := gst.NewElement("freeverb")
	if err != nil {
		return nil, err
	}
	return &GstFreeverb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFreeverbWithName(name string) (*GstFreeverb, error) {
	e, err := gst.NewElementWithName("freeverb", name)
	if err != nil {
		return nil, err
	}
	return &GstFreeverb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Stereo panorama width
// Default: 1, Min: 0, Max: 1
func (e *GstFreeverb) SetWidth(value float32) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstFreeverb) GetWidth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Damping of high frequencies
// Default: 0.2, Min: 0, Max: 1
func (e *GstFreeverb) SetDamping(value float32) error {
	return e.Element.SetProperty("damping", value)
}

func (e *GstFreeverb) GetDamping() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("damping")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// dry/wet level
// Default: 0.5, Min: 0, Max: 1
func (e *GstFreeverb) SetLevel(value float32) error {
	return e.Element.SetProperty("level", value)
}

func (e *GstFreeverb) GetLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Size of the simulated room
// Default: 0.5, Min: 0, Max: 1
func (e *GstFreeverb) SetRoomSize(value float32) error {
	return e.Element.SetProperty("room-size", value)
}

func (e *GstFreeverb) GetRoomSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("room-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}
