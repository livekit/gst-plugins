// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A standard X based videosink
type GstXImageSink struct {
	*GstVideoSink
}

func NewXImageSink() (*GstXImageSink, error) {
	e, err := gst.NewElement("ximagesink")
	if err != nil {
		return nil, err
	}
	return &GstXImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewXImageSinkWithName(name string) (*GstXImageSink, error) {
	e, err := gst.NewElementWithName("ximagesink", name)
	if err != nil {
		return nil, err
	}
	return &GstXImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// When enabled, runs the X display in synchronous mode. (unrelated to A/V sync, used only for debugging)
// Default: false
func (e *GstXImageSink) SetSynchronous(value bool) error {
	return e.Element.SetProperty("synchronous", value)
}

func (e *GstXImageSink) GetSynchronous() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synchronous")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Height of the window
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstXImageSink) GetWindowHeight() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Width of the window
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstXImageSink) GetWindowWidth() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// X Display name
// Default: NULL
func (e *GstXImageSink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstXImageSink) GetDisplay() (string, error) {
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

// When enabled, reverse caps negotiation (scaling) will respect original aspect ratio
// Default: true
func (e *GstXImageSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstXImageSink) GetForceAspectRatio() (bool, error) {
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

// When enabled, XEvents will be selected and handled
// Default: true
func (e *GstXImageSink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

func (e *GstXImageSink) GetHandleEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, the current frame will always be drawn in response to X Expose events
// Default: true
func (e *GstXImageSink) SetHandleExpose(value bool) error {
	return e.Element.SetProperty("handle-expose", value)
}

func (e *GstXImageSink) GetHandleExpose() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-expose")
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
// Default: NULL
func (e *GstXImageSink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstXImageSink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
