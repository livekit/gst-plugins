// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Display data using a Direct3D9 video renderer
type GstD3DVideoSink struct {
	*GstVideoSink
}

func NewD3DVideoSink() (*GstD3DVideoSink, error) {
	e, err := gst.NewElement("d3dvideosink")
	if err != nil {
		return nil, err
	}
	return &GstD3DVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewD3DVideoSinkWithName(name string) (*GstD3DVideoSink, error) {
	e, err := gst.NewElementWithName("d3dvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstD3DVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// If no window ID is given, a new render window is created
// Default: true
func (e *GstD3DVideoSink) SetCreateRenderWindow(value bool) error {
	return e.Element.SetProperty("create-render-window", value)
}

func (e *GstD3DVideoSink) GetCreateRenderWindow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-render-window")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, navigation events are sent upstream
// Default: true
func (e *GstD3DVideoSink) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

func (e *GstD3DVideoSink) GetEnableNavigationEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-navigation-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstD3DVideoSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstD3DVideoSink) GetForceAspectRatio() (bool, error) {
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

// If the render window is closed stop stream
// Default: true
func (e *GstD3DVideoSink) SetStreamStopOnClose(value bool) error {
	return e.Element.SetProperty("stream-stop-on-close", value)
}

func (e *GstD3DVideoSink) GetStreamStopOnClose() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("stream-stop-on-close")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
