// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A Xv based videosink
type GstXvImageSink struct {
	*GstVideoSink
}

func NewXvImageSink() (*GstXvImageSink, error) {
	e, err := gst.NewElement("xvimagesink")
	if err != nil {
		return nil, err
	}
	return &GstXvImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewXvImageSinkWithName(name string) (*GstXvImageSink, error) {
	e, err := gst.NewElementWithName("xvimagesink", name)
	if err != nil {
		return nil, err
	}
	return &GstXvImageSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// When enabled, the current frame will always be drawn in response to X Expose events
// Default: true
func (e *GstXvImageSink) SetHandleExpose(value bool) error {
	return e.Element.SetProperty("handle-expose", value)
}

func (e *GstXvImageSink) GetHandleExpose() (bool, error) {
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

// The render rectangle ('<x, y, width, height>')

func (e *GstXvImageSink) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// The saturation of the video
// Default: 0, Min: -1000, Max: 1000
func (e *GstXvImageSink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstXvImageSink) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether to autofill overlay with colorkey
// Default: true
func (e *GstXvImageSink) SetAutopaintColorkey(value bool) error {
	return e.Element.SetProperty("autopaint-colorkey", value)
}

func (e *GstXvImageSink) GetAutopaintColorkey() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autopaint-colorkey")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The name of the video adaptor
// Default: NULL
func (e *GstXvImageSink) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// When enabled, XEvents will be selected and handled
// Default: true
func (e *GstXvImageSink) SetHandleEvents(value bool) error {
	return e.Element.SetProperty("handle-events", value)
}

func (e *GstXvImageSink) GetHandleEvents() (bool, error) {
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

// The hue of the video
// Default: 0, Min: -1000, Max: 1000
func (e *GstXvImageSink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstXvImageSink) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of the window
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstXvImageSink) GetWindowWidth() (uint64, error) {
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

// The brightness of the video
// Default: 0, Min: -1000, Max: 1000
func (e *GstXvImageSink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstXvImageSink) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Color to use for the overlay mask
// Default: 526352, Min: -2147483648, Max: 2147483647
func (e *GstXvImageSink) SetColorkey(value int) error {
	return e.Element.SetProperty("colorkey", value)
}

func (e *GstXvImageSink) GetColorkey() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorkey")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstXvImageSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstXvImageSink) GetForceAspectRatio() (bool, error) {
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

// The number of the video adaptor
// Default: 0
func (e *GstXvImageSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstXvImageSink) GetDevice() (string, error) {
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

// Draw black borders to fill unused area in force-aspect-ratio mode
// Default: true
func (e *GstXvImageSink) SetDrawBorders(value bool) error {
	return e.Element.SetProperty("draw-borders", value)
}

func (e *GstXvImageSink) GetDrawBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to double-buffer the output
// Default: false
func (e *GstXvImageSink) SetDoubleBuffer(value bool) error {
	return e.Element.SetProperty("double-buffer", value)
}

func (e *GstXvImageSink) GetDoubleBuffer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("double-buffer")
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
func (e *GstXvImageSink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

func (e *GstXvImageSink) GetPixelAspectRatio() (string, error) {
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

// When enabled, runs the X display in synchronous mode. (unrelated to A/V sync, used only for debugging)
// Default: false
func (e *GstXvImageSink) SetSynchronous(value bool) error {
	return e.Element.SetProperty("synchronous", value)
}

func (e *GstXvImageSink) GetSynchronous() (bool, error) {
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
func (e *GstXvImageSink) GetWindowHeight() (uint64, error) {
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

// The contrast of the video
// Default: 0, Min: -1000, Max: 1000
func (e *GstXvImageSink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstXvImageSink) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// X Display name
// Default: NULL
func (e *GstXvImageSink) SetDisplay(value string) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstXvImageSink) GetDisplay() (string, error) {
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
