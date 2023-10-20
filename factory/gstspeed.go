// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Set speed/pitch on audio/raw streams (resampler)
type GstSpeed struct {
	*gst.Element
}

func NewSpeed() (*GstSpeed, error) {
	e, err := gst.NewElement("speed")
	if err != nil {
		return nil, err
	}
	return &GstSpeed{Element: e}, nil
}

func NewSpeedWithName(name string) (*GstSpeed, error) {
	e, err := gst.NewElementWithName("speed", name)
	if err != nil {
		return nil, err
	}
	return &GstSpeed{Element: e}, nil
}

// speed
// Default: 1, Min: 0.1, Max: 40
func (e *GstSpeed) SetSpeed(value float32) error {
	return e.Element.SetProperty("speed", value)
}

func (e *GstSpeed) GetSpeed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}
