// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Video sink using the Linux kernel mode setting API
type GstKMSSink struct {
	*GstVideoSink
}

func NewKMSSink() (*GstKMSSink, error) {
	e, err := gst.NewElement("kmssink")
	if err != nil {
		return nil, err
	}
	return &GstKMSSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewKMSSinkWithName(name string) (*GstKMSSink, error) {
	e, err := gst.NewElementWithName("kmssink", name)
	if err != nil {
		return nil, err
	}
	return &GstKMSSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Height of the display surface in pixels
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKMSSink) GetDisplayHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("display-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of the display surface in pixels
// Default: 0, Min: 0, Max: 2147483647
func (e *GstKMSSink) GetDisplayWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("display-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// DRM device driver name
// Default: NULL
func (e *GstKMSSink) SetDriverName(value string) error {
	return e.Element.SetProperty("driver-name", value)
}

func (e *GstKMSSink) GetDriverName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("driver-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// DRM file descriptor
// Default: -1, Min: -1, Max: 2147483647
func (e *GstKMSSink) SetFd(value int) error {
	return e.Element.SetProperty("fd", value)
}

func (e *GstKMSSink) GetFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// DRM connector id
// Default: -1, Min: -1, Max: 2147483647
func (e *GstKMSSink) SetConnectorId(value int) error {
	return e.Element.SetProperty("connector-id", value)
}

func (e *GstKMSSink) GetConnectorId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("connector-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Additional properties for the plane

func (e *GstKMSSink) SetPlaneProperties(value *gst.Structure) error {
	return e.Element.SetProperty("plane-properties", value)
}

func (e *GstKMSSink) GetPlaneProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("plane-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// DRM bus ID
// Default: NULL
func (e *GstKMSSink) SetBusId(value string) error {
	return e.Element.SetProperty("bus-id", value)
}

func (e *GstKMSSink) GetBusId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bus-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Additional properties for the connector

func (e *GstKMSSink) SetConnectorProperties(value *gst.Structure) error {
	return e.Element.SetProperty("connector-properties", value)
}

func (e *GstKMSSink) GetConnectorProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("connector-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// When enabled, the sink try to configure the display mode
// Default: false
func (e *GstKMSSink) SetForceModesetting(value bool) error {
	return e.Element.SetProperty("force-modesetting", value)
}

func (e *GstKMSSink) GetForceModesetting() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-modesetting")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled and CRTC was set with a new mode, previous CRTC mode willbe restored when going to NULL state.
// Default: true
func (e *GstKMSSink) SetRestoreCrtc(value bool) error {
	return e.Element.SetProperty("restore-crtc", value)
}

func (e *GstKMSSink) GetRestoreCrtc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("restore-crtc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled will not wait internally for vsync. Should be used for atomic drivers to avoid double vsync.
// Default: false
func (e *GstKMSSink) SetSkipVsync(value bool) error {
	return e.Element.SetProperty("skip-vsync", value)
}

func (e *GstKMSSink) GetSkipVsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-vsync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// User can tell kmssink if the driver can support scale
// Default: true
func (e *GstKMSSink) SetCanScale(value bool) error {
	return e.Element.SetProperty("can-scale", value)
}

func (e *GstKMSSink) GetCanScale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DRM plane id
// Default: -1, Min: -1, Max: 2147483647
func (e *GstKMSSink) SetPlaneId(value int) error {
	return e.Element.SetProperty("plane-id", value)
}

func (e *GstKMSSink) GetPlaneId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("plane-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The render rectangle ('<x, y, width, height>')

func (e *GstKMSSink) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}
