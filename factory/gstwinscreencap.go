// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Captures screen
type GstDX9ScreenCapSrc struct {
	*base.GstPushSrc
}

func NewDX9ScreenCapSrc() (*GstDX9ScreenCapSrc, error) {
	e, err := gst.NewElement("dx9screencapsrc")
	if err != nil {
		return nil, err
	}
	return &GstDX9ScreenCapSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDX9ScreenCapSrcWithName(name string) (*GstDX9ScreenCapSrc, error) {
	e, err := gst.NewElementWithName("dx9screencapsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDX9ScreenCapSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Horizontal coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDX9ScreenCapSrc) SetX(value int) error {
	return e.Element.SetProperty("x", value)
}

func (e *GstDX9ScreenCapSrc) GetX() (int, error) {
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

// Vertical coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDX9ScreenCapSrc) SetY(value int) error {
	return e.Element.SetProperty("y", value)
}

func (e *GstDX9ScreenCapSrc) GetY() (int, error) {
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

// Whether to show mouse cursor (default off)
// Default: false
func (e *GstDX9ScreenCapSrc) SetCursor(value bool) error {
	return e.Element.SetProperty("cursor", value)
}

func (e *GstDX9ScreenCapSrc) GetCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Height of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDX9ScreenCapSrc) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstDX9ScreenCapSrc) GetHeight() (int, error) {
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

// Which monitor to use (0 = 1st monitor and default)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDX9ScreenCapSrc) SetMonitor(value int) error {
	return e.Element.SetProperty("monitor", value)
}

func (e *GstDX9ScreenCapSrc) GetMonitor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("monitor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDX9ScreenCapSrc) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstDX9ScreenCapSrc) GetWidth() (int, error) {
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

// Captures screen
type GstGDIScreenCapSrc struct {
	*base.GstPushSrc
}

func NewGDIScreenCapSrc() (*GstGDIScreenCapSrc, error) {
	e, err := gst.NewElement("gdiscreencapsrc")
	if err != nil {
		return nil, err
	}
	return &GstGDIScreenCapSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewGDIScreenCapSrcWithName(name string) (*GstGDIScreenCapSrc, error) {
	e, err := gst.NewElementWithName("gdiscreencapsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstGDIScreenCapSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Whether to show mouse cursor (default off)
// Default: false
func (e *GstGDIScreenCapSrc) SetCursor(value bool) error {
	return e.Element.SetProperty("cursor", value)
}

func (e *GstGDIScreenCapSrc) GetCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Height of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGDIScreenCapSrc) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstGDIScreenCapSrc) GetHeight() (int, error) {
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

// Which monitor to use (0 = 1st monitor and default)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGDIScreenCapSrc) SetMonitor(value int) error {
	return e.Element.SetProperty("monitor", value)
}

func (e *GstGDIScreenCapSrc) GetMonitor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("monitor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGDIScreenCapSrc) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstGDIScreenCapSrc) GetWidth() (int, error) {
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

// Horizontal coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGDIScreenCapSrc) SetX(value int) error {
	return e.Element.SetProperty("x", value)
}

func (e *GstGDIScreenCapSrc) GetX() (int, error) {
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

// Vertical coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGDIScreenCapSrc) SetY(value int) error {
	return e.Element.SetProperty("y", value)
}

func (e *GstGDIScreenCapSrc) GetY() (int, error) {
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
