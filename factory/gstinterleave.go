// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Splits one interleaved multichannel audio stream into many mono audio streams
type GstDeinterleave struct {
	*gst.Element
}

func NewDeinterleave() (*GstDeinterleave, error) {
	e, err := gst.NewElement("deinterleave")
	if err != nil {
		return nil, err
	}
	return &GstDeinterleave{Element: e}, nil
}

func NewDeinterleaveWithName(name string) (*GstDeinterleave, error) {
	e, err := gst.NewElementWithName("deinterleave", name)
	if err != nil {
		return nil, err
	}
	return &GstDeinterleave{Element: e}, nil
}

// Keep the original channel positions on the output buffers
// Default: false
func (e *GstDeinterleave) SetKeepPositions(value bool) error {
	return e.Element.SetProperty("keep-positions", value)
}

func (e *GstDeinterleave) GetKeepPositions() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-positions")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Folds many mono channels into one interleaved audio stream
type GstInterleave struct {
	*gst.Element
}

func NewInterleave() (*GstInterleave, error) {
	e, err := gst.NewElement("interleave")
	if err != nil {
		return nil, err
	}
	return &GstInterleave{Element: e}, nil
}

func NewInterleaveWithName(name string) (*GstInterleave, error) {
	e, err := gst.NewElementWithName("interleave", name)
	if err != nil {
		return nil, err
	}
	return &GstInterleave{Element: e}, nil
}

// Take channel positions from the input
// Default: true
func (e *GstInterleave) SetChannelPositionsFromInput(value bool) error {
	return e.Element.SetProperty("channel-positions-from-input", value)
}

func (e *GstInterleave) GetChannelPositionsFromInput() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("channel-positions-from-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Channel positions used on the output

func (e *GstInterleave) SetChannelPositions(value interface{}) error {
	return e.Element.SetProperty("channel-positions", value)
}

func (e *GstInterleave) GetChannelPositions() (interface{}, error) {
	return e.Element.GetProperty("channel-positions")
}
