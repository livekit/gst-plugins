// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A video sink that renders to a GtkWidget using OpenGL
type GstGtkGLSink struct {
	*GstGtkBaseSink
}

func NewGtkGLSink() (*GstGtkGLSink, error) {
	e, err := gst.NewElement("gtkglsink")
	if err != nil {
		return nil, err
	}
	return &GstGtkGLSink{GstGtkBaseSink: &GstGtkBaseSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewGtkGLSinkWithName(name string) (*GstGtkGLSink, error) {
	e, err := gst.NewElementWithName("gtkglsink", name)
	if err != nil {
		return nil, err
	}
	return &GstGtkGLSink{GstGtkBaseSink: &GstGtkBaseSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// rotate method
// Default: identity (0)
func (e *GstGtkGLSink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

func (e *GstGtkGLSink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

// A video sink that renders to a GtkWidget
type GstGtkSink struct {
	*GstGtkBaseSink
}

func NewGtkSink() (*GstGtkSink, error) {
	e, err := gst.NewElement("gtksink")
	if err != nil {
		return nil, err
	}
	return &GstGtkSink{GstGtkBaseSink: &GstGtkBaseSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewGtkSinkWithName(name string) (*GstGtkSink, error) {
	e, err := gst.NewElementWithName("gtksink", name)
	if err != nil {
		return nil, err
	}
	return &GstGtkSink{GstGtkBaseSink: &GstGtkBaseSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

type GstGtkBaseSink struct {
	*GstVideoSink
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstGtkBaseSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstGtkBaseSink) GetForceAspectRatio() (bool, error) {
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

// When enabled, alpha will be ignored and converted to black
// Default: true
func (e *GstGtkBaseSink) SetIgnoreAlpha(value bool) error {
	return e.Element.SetProperty("ignore-alpha", value)
}

func (e *GstGtkBaseSink) GetIgnoreAlpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The pixel aspect ratio of the device
// Default: 0/1, Min: 0/2147483647, Max: 2147483647/1
func (e *GstGtkBaseSink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstGtkBaseSink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

// The pixel aspect ratio of the video (0/1 = follow stream)
// Default: 0/1, Min: 0/2147483647, Max: 2147483647/1
func (e *GstGtkBaseSink) SetVideoAspectRatioOverride(value interface{}) error {
	return e.Element.SetProperty("video-aspect-ratio-override", value)
}

func (e *GstGtkBaseSink) GetVideoAspectRatioOverride() (interface{}, error) {
	return e.Element.GetProperty("video-aspect-ratio-override")
}

// The GtkWidget to place in the widget hierarchy (must only be get from the GTK main thread)

func (e *GstGtkBaseSink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}
