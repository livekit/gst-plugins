// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A video sink that renders to a GtkWidget using Wayland API
type GstGtkWaylandSink struct {
	*GstVideoSink
}

func NewGtkWaylandSinkWithProperties(properties map[string]interface{}) (*GstGtkWaylandSink, error) {
	e, err := gst.NewElementWithProperties("gtkwaylandsink", properties)
	if err != nil {
		return nil, err
	}
	return &GstGtkWaylandSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Path of the DRM device to use for dumb buffer allocation
// Default: NULL
func (e *GstGtkWaylandSink) SetDrmDevice(value string) error {
	return e.Element.SetProperty("drm-device", value)
}

func (e *GstGtkWaylandSink) GetDrmDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("drm-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// rotate method
// Default: identity (0)
func (e *GstGtkWaylandSink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

func (e *GstGtkWaylandSink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

// The GtkWidget to place in the widget hierarchy (must only be get from the GTK main thread)

func (e *GstGtkWaylandSink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}
