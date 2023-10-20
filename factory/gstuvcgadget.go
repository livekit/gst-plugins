// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Streams Video via UVC Gadget
type GstUvcSink struct {
	*gst.Bin
}

func NewUvcSink() (*GstUvcSink, error) {
	e, err := gst.NewElement("uvcsink")
	if err != nil {
		return nil, err
	}
	return &GstUvcSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewUvcSinkWithName(name string) (*GstUvcSink, error) {
	e, err := gst.NewElementWithName("uvcsink", name)
	if err != nil {
		return nil, err
	}
	return &GstUvcSink{Bin: &gst.Bin{Element: e}}, nil
}

// The stream status of the host
// Default: false
func (e *GstUvcSink) GetStreaming() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streaming")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
