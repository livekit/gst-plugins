// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates audio test signals identifying channels
type GstFliteTestSrc struct {
	*base.GstBaseSrc
}

func NewFliteTestSrc() (*GstFliteTestSrc, error) {
	e, err := gst.NewElement("flitetestsrc")
	if err != nil {
		return nil, err
	}
	return &GstFliteTestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewFliteTestSrcWithName(name string) (*GstFliteTestSrc, error) {
	e, err := gst.NewElementWithName("flitetestsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstFliteTestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Number of samples in each outgoing buffer
// Default: 1024, Min: 1, Max: 2147483647
func (e *GstFliteTestSrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

func (e *GstFliteTestSrc) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
