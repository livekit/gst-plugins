// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Provides various Image Quality Assessment metrics
type GstIqa struct {
	*GstVideoAggregator
}

func NewIqa() (*GstIqa, error) {
	e, err := gst.NewElement("iqa")
	if err != nil {
		return nil, err
	}
	return &GstIqa{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewIqaWithName(name string) (*GstIqa, error) {
	e, err := gst.NewElementWithName("iqa", name)
	if err != nil {
		return nil, err
	}
	return &GstIqa{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// dssim value over which the element will post an error message on the bus. A value < 0.0 means 'disabled'.
// Default: 0, Min: -1, Max: 1.79769e+308
func (e *GstIqa) SetDssimErrorThreshold(value float64) error {
	return e.Element.SetProperty("dssim-error-threshold", value)
}

func (e *GstIqa) GetDssimErrorThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("dssim-error-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Controls the frame comparison mode.
// Default: (none)
func (e *GstIqa) SetMode(value GstIqaMode) error {
	e.Element.SetArg("mode", fmt.Sprint(value))
	return nil
}

func (e *GstIqa) GetMode() (GstIqaMode, error) {
	var value GstIqaMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstIqaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstIqaMode")
	}
	return value, nil
}

// Run structural similarity checks
// Default: false
func (e *GstIqa) SetDoDssim(value bool) error {
	return e.Element.SetProperty("do-dssim", value)
}

func (e *GstIqa) GetDoDssim() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-dssim")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstIqaMode int

const (
	GstIqaModeStrict GstIqaMode = 0x00000002 // Strict comparison of frames.
)
