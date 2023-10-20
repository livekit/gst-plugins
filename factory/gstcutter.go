// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Audio Cutter to split audio into non-silent bits
type GstCutter struct {
	*gst.Element
}

func NewCutter() (*GstCutter, error) {
	e, err := gst.NewElement("cutter")
	if err != nil {
		return nil, err
	}
	return &GstCutter{Element: e}, nil
}

func NewCutterWithName(name string) (*GstCutter, error) {
	e, err := gst.NewElementWithName("cutter", name)
	if err != nil {
		return nil, err
	}
	return &GstCutter{Element: e}, nil
}

// Volume threshold before trigger (in dB)
// Default: -46.0517, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstCutter) SetThresholdDB(value float64) error {
	return e.Element.SetProperty("threshold-dB", value)
}

func (e *GstCutter) GetThresholdDB() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold-dB")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// do we leak buffers when below threshold ?
// Default: false
func (e *GstCutter) SetLeaky(value bool) error {
	return e.Element.SetProperty("leaky", value)
}

func (e *GstCutter) GetLeaky() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("leaky")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Length of pre-recording buffer (in nanoseconds)
// Default: 200000000, Min: 0, Max: 18446744073709551615
func (e *GstCutter) SetPreLength(value uint64) error {
	return e.Element.SetProperty("pre-length", value)
}

func (e *GstCutter) GetPreLength() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("pre-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Length of drop below threshold before cut_stop (in nanoseconds)
// Default: 500000000, Min: 0, Max: 18446744073709551615
func (e *GstCutter) SetRunLength(value uint64) error {
	return e.Element.SetProperty("run-length", value)
}

func (e *GstCutter) GetRunLength() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("run-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Volume threshold before trigger
// Default: 0.1, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstCutter) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstCutter) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}
