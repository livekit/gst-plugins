// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Linux framebuffer videosink
type GstFBDEVSink struct {
	*GstVideoSink
}

func NewFBDEVSink() (*GstFBDEVSink, error) {
	e, err := gst.NewElement("fbdevsink")
	if err != nil {
		return nil, err
	}
	return &GstFBDEVSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewFBDEVSinkWithName(name string) (*GstFBDEVSink, error) {
	e, err := gst.NewElementWithName("fbdevsink", name)
	if err != nil {
		return nil, err
	}
	return &GstFBDEVSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// The framebuffer device eg: /dev/fb0
// Default: NULL
func (e *GstFBDEVSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstFBDEVSink) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
