// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output to wayland surface
type GstWaylandSink struct {
	*GstVideoSink
}

func NewWaylandSinkWithProperties(properties map[string]interface{}) (*GstWaylandSink, error) {
	e, err := gst.NewElementWithProperties("waylandsink", properties)
	if err != nil {
		return nil, err
	}
	return &GstWaylandSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Wayland display name to connect to, if not supplied via the GstContext
// Default: NULL
func (e *GstWaylandSink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstWaylandSink) GetDisplay() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Path of the DRM device to use for dumb buffer allocation
// Default: NULL
func (e *GstWaylandSink) SetDrmDevice(value string) error {
	return e.Element.SetProperty("drm-device", value)
}

func (e *GstWaylandSink) GetDrmDevice() (string, error) {
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

// Whether the surface should be made fullscreen
// Default: false
func (e *GstWaylandSink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

func (e *GstWaylandSink) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The render rectangle ('<x, y, width, height>')

func (e *GstWaylandSink) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// rotate method
// Default: identity (0)
func (e *GstWaylandSink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

func (e *GstWaylandSink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}
