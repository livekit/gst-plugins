// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Decodes images in a video stream using GdkPixbuf
type GstGdkPixbufDec struct {
	*gst.Element
}

func NewGdkPixbufDec() (*GstGdkPixbufDec, error) {
	e, err := gst.NewElement("gdkpixbufdec")
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufDec{Element: e}, nil
}

func NewGdkPixbufDecWithName(name string) (*GstGdkPixbufDec, error) {
	e, err := gst.NewElementWithName("gdkpixbufdec", name)
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufDec{Element: e}, nil
}

// Overlay an image onto a video stream
type GstGdkPixbufOverlay struct {
	*GstVideoFilter
}

func NewGdkPixbufOverlay() (*GstGdkPixbufOverlay, error) {
	e, err := gst.NewElement("gdkpixbufoverlay")
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufOverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewGdkPixbufOverlayWithName(name string) (*GstGdkPixbufOverlay, error) {
	e, err := gst.NewElementWithName("gdkpixbufoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufOverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Positioning mode of offset-x and offset-y properties
// Default: pixels-relative-to-edges (0)
func (e *GstGdkPixbufOverlay) SetPositioningMode(value GstGdkPixbufPositioningMode) error {
	e.Element.SetArg("positioning-mode", string(value))
	return nil
}

func (e *GstGdkPixbufOverlay) GetPositioningMode() (GstGdkPixbufPositioningMode, error) {
	var value GstGdkPixbufPositioningMode
	var ok bool
	v, err := e.Element.GetProperty("positioning-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGdkPixbufPositioningMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGdkPixbufPositioningMode")
	}
	return value, nil
}

// Vertical offset of overlay image in fractions of video image height, from top-left corner of video image (in relative positioning)
// Default: 0, Min: -1, Max: 1
func (e *GstGdkPixbufOverlay) SetRelativeY(value float64) error {
	return e.Element.SetProperty("relative-y", value)
}

func (e *GstGdkPixbufOverlay) GetRelativeY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Global alpha of overlay image
// Default: 1, Min: 0, Max: 1
func (e *GstGdkPixbufOverlay) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstGdkPixbufOverlay) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Horizontal offset of overlay image in fractions of video image width, from top-left corner of video image (absolute positioning)
// Default: 0, Min: -1, Max: 1
func (e *GstGdkPixbufOverlay) SetCoefX(value float64) error {
	return e.Element.SetProperty("coef-x", value)
}

func (e *GstGdkPixbufOverlay) GetCoefX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("coef-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Vertical offset of overlay image in fractions of video image height, from top-left corner of video image (absolute positioning)
// Default: 0, Min: -1, Max: 1
func (e *GstGdkPixbufOverlay) SetCoefY(value float64) error {
	return e.Element.SetProperty("coef-y", value)
}

func (e *GstGdkPixbufOverlay) GetCoefY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("coef-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Location of image file to overlay
// Default: NULL
func (e *GstGdkPixbufOverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGdkPixbufOverlay) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Width of overlay image in pixels (0 = same as overlay image)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGdkPixbufOverlay) SetOverlayWidth(value int) error {
	return e.Element.SetProperty("overlay-width", value)
}

func (e *GstGdkPixbufOverlay) GetOverlayWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// For positive value, horizontal offset of overlay image in pixels from left of video image. For negative value, horizontal offset of overlay image in pixels from right of video image
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstGdkPixbufOverlay) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

func (e *GstGdkPixbufOverlay) GetOffsetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// For positive value, vertical offset of overlay image in pixels from top of video image. For negative value, vertical offset of overlay image in pixels from bottom of video image
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstGdkPixbufOverlay) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

func (e *GstGdkPixbufOverlay) GetOffsetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Height of overlay image in pixels (0 = same as overlay image)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGdkPixbufOverlay) SetOverlayHeight(value int) error {
	return e.Element.SetProperty("overlay-height", value)
}

func (e *GstGdkPixbufOverlay) GetOverlayHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// GdkPixbuf object to render

func (e *GstGdkPixbufOverlay) SetPixbuf(value interface{}) error {
	return e.Element.SetProperty("pixbuf", value)
}

func (e *GstGdkPixbufOverlay) GetPixbuf() (interface{}, error) {
	return e.Element.GetProperty("pixbuf")
}

// Horizontal offset of overlay image in fractions of video image width, from top-left corner of video image (in relative positioning)
// Default: 0, Min: -1, Max: 1
func (e *GstGdkPixbufOverlay) SetRelativeX(value float64) error {
	return e.Element.SetProperty("relative-x", value)
}

func (e *GstGdkPixbufOverlay) GetRelativeX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("relative-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Output images as GdkPixbuf objects in bus messages
type GstGdkPixbufSink struct {
	*GstVideoSink
}

func NewGdkPixbufSink() (*GstGdkPixbufSink, error) {
	e, err := gst.NewElement("gdkpixbufsink")
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewGdkPixbufSinkWithName(name string) (*GstGdkPixbufSink, error) {
	e, err := gst.NewElementWithName("gdkpixbufsink", name)
	if err != nil {
		return nil, err
	}
	return &GstGdkPixbufSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Last GdkPixbuf object rendered

func (e *GstGdkPixbufSink) GetLastPixbuf() (interface{}, error) {
	return e.Element.GetProperty("last-pixbuf")
}

// Whether to post messages containing pixbufs on the bus
// Default: true
func (e *GstGdkPixbufSink) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstGdkPixbufSink) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstGdkPixbufPositioningMode string

const (
	GstGdkPixbufPositioningModePixelsRelativeToEdges GstGdkPixbufPositioningMode = "pixels-relative-to-edges" // pixels-relative-to-edges
	GstGdkPixbufPositioningModePixelsAbsolute        GstGdkPixbufPositioningMode = "pixels-absolute"          // pixels-absolute
)
