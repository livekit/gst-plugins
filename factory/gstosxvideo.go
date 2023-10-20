// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// OSX native videosink
type GstOSXVideoSink struct {
	*GstVideoSink
}

func NewOSXVideoSink() (*GstOSXVideoSink, error) {
	e, err := gst.NewElement("osxvideosink")
	if err != nil {
		return nil, err
	}
	return &GstOSXVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewOSXVideoSinkWithName(name string) (*GstOSXVideoSink, error) {
	e, err := gst.NewElementWithName("osxvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstOSXVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// For ABI compatibility only, do not use
// Default: false
func (e *GstOSXVideoSink) SetEmbed(value bool) error {
	return e.Element.SetProperty("embed", value)
}

func (e *GstOSXVideoSink) GetEmbed() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("embed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ration
// Default: false
func (e *GstOSXVideoSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstOSXVideoSink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
