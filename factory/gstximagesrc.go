// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates a screenshot video stream
type GstXImageSrc struct {
	*base.GstPushSrc
}

func NewXImageSrc() (*GstXImageSrc, error) {
	e, err := gst.NewElement("ximagesrc")
	if err != nil {
		return nil, err
	}
	return &GstXImageSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewXImageSrcWithName(name string) (*GstXImageSrc, error) {
	e, err := gst.NewElementWithName("ximagesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstXImageSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Show mouse pointer (if XFixes extension enabled)
// Default: true
func (e *GstXImageSrc) SetShowPointer(value bool) error {
	return e.Element.SetProperty("show-pointer", value)
}

func (e *GstXImageSrc) GetShowPointer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-pointer")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// X coordinate of top left corner of area to be recorded (0 for top left of screen)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstXImageSrc) SetStartx(value uint) error {
	return e.Element.SetProperty("startx", value)
}

func (e *GstXImageSrc) GetStartx() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("startx")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// X Display Name
// Default: NULL
func (e *GstXImageSrc) SetDisplayName(value string) error {
	return e.Element.SetProperty("display-name", value)
}

func (e *GstXImageSrc) GetDisplayName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// When enabled, navigation events are handled
// Default: false
func (e *GstXImageSrc) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

func (e *GstXImageSrc) GetEnableNavigationEvents() (bool, error) {
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

// Whether the display is remote
// Default: false
func (e *GstXImageSrc) SetRemote(value bool) error {
	return e.Element.SetProperty("remote", value)
}

func (e *GstXImageSrc) GetRemote() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remote")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use XDamage (if XDamage extension enabled)
// Default: true
func (e *GstXImageSrc) SetUseDamage(value bool) error {
	return e.Element.SetProperty("use-damage", value)
}

func (e *GstXImageSrc) GetUseDamage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-damage")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Window XID to capture from
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstXImageSrc) SetXid(value uint64) error {
	return e.Element.SetProperty("xid", value)
}

func (e *GstXImageSrc) GetXid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("xid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Window name to capture from
// Default: NULL
func (e *GstXImageSrc) SetXname(value string) error {
	return e.Element.SetProperty("xname", value)
}

func (e *GstXImageSrc) GetXname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("xname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// X coordinate of bottom right corner of area to be recorded (0 for bottom right of screen)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstXImageSrc) SetEndx(value uint) error {
	return e.Element.SetProperty("endx", value)
}

func (e *GstXImageSrc) GetEndx() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("endx")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Y coordinate of bottom right corner of area to be recorded (0 for bottom right of screen)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstXImageSrc) SetEndy(value uint) error {
	return e.Element.SetProperty("endy", value)
}

func (e *GstXImageSrc) GetEndy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("endy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Y coordinate of top left corner of area to be recorded (0 for top left of screen)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstXImageSrc) SetStarty(value uint) error {
	return e.Element.SetProperty("starty", value)
}

func (e *GstXImageSrc) GetStarty() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("starty")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
