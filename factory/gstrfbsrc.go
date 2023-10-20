// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates a rfb video stream
type GstRfbSrc struct {
	*base.GstPushSrc
}

func NewRfbSrc() (*GstRfbSrc, error) {
	e, err := gst.NewElement("rfbsrc")
	if err != nil {
		return nil, err
	}
	return &GstRfbSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewRfbSrcWithName(name string) (*GstRfbSrc, error) {
	e, err := gst.NewElementWithName("rfbsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRfbSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Incremental updates
// Default: true
func (e *GstRfbSrc) SetIncremental(value bool) error {
	return e.Element.SetProperty("incremental", value)
}

func (e *GstRfbSrc) GetIncremental() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("incremental")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// y offset for screen scrapping
// Default: 0, Min: 0, Max: 65535
func (e *GstRfbSrc) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

func (e *GstRfbSrc) GetOffsetY() (int, error) {
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

// Password for authentication

func (e *GstRfbSrc) SetPassword(value string) error {
	return e.Element.SetProperty("password", value)
}

// Port
// Default: 5900, Min: 1, Max: 65535
func (e *GstRfbSrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstRfbSrc) GetPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use copyrect encoding
// Default: false
func (e *GstRfbSrc) SetUseCopyrect(value bool) error {
	return e.Element.SetProperty("use-copyrect", value)
}

func (e *GstRfbSrc) GetUseCopyrect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-copyrect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// RFB protocol version
// Default: 3.3
func (e *GstRfbSrc) SetVersion(value string) error {
	return e.Element.SetProperty("version", value)
}

func (e *GstRfbSrc) GetVersion() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("version")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// only view the desktop
// Default: false
func (e *GstRfbSrc) SetViewOnly(value bool) error {
	return e.Element.SetProperty("view-only", value)
}

func (e *GstRfbSrc) GetViewOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("view-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// width of screen
// Default: 0, Min: 0, Max: 65535
func (e *GstRfbSrc) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstRfbSrc) GetWidth() (int, error) {
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

// height of screen
// Default: 0, Min: 0, Max: 65535
func (e *GstRfbSrc) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstRfbSrc) GetHeight() (int, error) {
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

// Host to connect to
// Default: 127.0.0.1
func (e *GstRfbSrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstRfbSrc) GetHost() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("host")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// x offset for screen scrapping
// Default: 0, Min: 0, Max: 65535
func (e *GstRfbSrc) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

func (e *GstRfbSrc) GetOffsetX() (int, error) {
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

// Share desktop with other clients
// Default: true
func (e *GstRfbSrc) SetShared(value bool) error {
	return e.Element.SetProperty("shared", value)
}

func (e *GstRfbSrc) GetShared() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("shared")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// URI in the form of rfb://host:port?query
// Default: rfb://127.0.0.1:5900
func (e *GstRfbSrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstRfbSrc) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
