// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A video sink that renders to a QQuickItem for Qt6 using Direct3D11
type GstQml6D3D11Sink struct {
	*GstVideoSink
}

func NewQml6D3D11Sink() (*GstQml6D3D11Sink, error) {
	e, err := gst.NewElement("qml6d3d11sink")
	if err != nil {
		return nil, err
	}
	return &GstQml6D3D11Sink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewQml6D3D11SinkWithName(name string) (*GstQml6D3D11Sink, error) {
	e, err := gst.NewElementWithName("qml6d3d11sink", name)
	if err != nil {
		return nil, err
	}
	return &GstQml6D3D11Sink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstQml6D3D11Sink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstQml6D3D11Sink) GetForceAspectRatio() (bool, error) {
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

// The QQuickItem to place in the object hierarchy

func (e *GstQml6D3D11Sink) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

func (e *GstQml6D3D11Sink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}
