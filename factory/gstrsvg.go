// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Uses librsvg to decode SVG images
type GstRsvgDec struct {
	*GstVideoDecoder
}

func NewRsvgDec() (*GstRsvgDec, error) {
	e, err := gst.NewElement("rsvgdec")
	if err != nil {
		return nil, err
	}
	return &GstRsvgDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewRsvgDecWithName(name string) (*GstRsvgDec, error) {
	e, err := gst.NewElementWithName("rsvgdec", name)
	if err != nil {
		return nil, err
	}
	return &GstRsvgDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Overlays SVG graphics over a video stream
type GstRsvgOverlay struct {
	*GstVideoFilter
}

func NewRsvgOverlay() (*GstRsvgOverlay, error) {
	e, err := gst.NewElement("rsvgoverlay")
	if err != nil {
		return nil, err
	}
	return &GstRsvgOverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewRsvgOverlayWithName(name string) (*GstRsvgOverlay, error) {
	e, err := gst.NewElementWithName("rsvgoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstRsvgOverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Specify a width relative to the display size.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstRsvgOverlay) SetWidthRelative(value float32) error {
	return e.Element.SetProperty("width-relative", value)
}

func (e *GstRsvgOverlay) GetWidthRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("width-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Specify an x offset.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstRsvgOverlay) SetX(value int) error {
	return e.Element.SetProperty("x", value)
}

func (e *GstRsvgOverlay) GetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify an x offset relative to the display size.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstRsvgOverlay) SetXRelative(value float32) error {
	return e.Element.SetProperty("x-relative", value)
}

func (e *GstRsvgOverlay) GetXRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("x-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Specify a y offset.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstRsvgOverlay) SetY(value int) error {
	return e.Element.SetProperty("y", value)
}

func (e *GstRsvgOverlay) GetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify a y offset relative to the display size.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstRsvgOverlay) SetYRelative(value float32) error {
	return e.Element.SetProperty("y-relative", value)
}

func (e *GstRsvgOverlay) GetYRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("y-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Specify a width in pixels.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstRsvgOverlay) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstRsvgOverlay) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// SVG data.

func (e *GstRsvgOverlay) SetData(value string) error {
	return e.Element.SetProperty("data", value)
}

// Fit the SVG to fill the whole frame.
// Default: false
func (e *GstRsvgOverlay) SetFitToFrame(value bool) error {
	return e.Element.SetProperty("fit-to-frame", value)
}

func (e *GstRsvgOverlay) GetFitToFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fit-to-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Specify a height in pixels.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstRsvgOverlay) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstRsvgOverlay) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify a height relative to the display size.
// Default: 0, Min: -3.40282e+38, Max: 3.40282e+38
func (e *GstRsvgOverlay) SetHeightRelative(value float32) error {
	return e.Element.SetProperty("height-relative", value)
}

func (e *GstRsvgOverlay) GetHeightRelative() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("height-relative")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// SVG file location.

func (e *GstRsvgOverlay) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}
