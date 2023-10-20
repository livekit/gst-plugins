// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Stream data from a video capture device through Windows kernel streaming
type GstKsVideoSrc struct {
	*base.GstPushSrc
}

func NewKsVideoSrc() (*GstKsVideoSrc, error) {
	e, err := gst.NewElement("ksvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstKsVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewKsVideoSrcWithName(name string) (*GstKsVideoSrc, error) {
	e, err := gst.NewElementWithName("ksvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstKsVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Last measured framerate, if statistics are enabled
// Default: -1, Min: -1, Max: 2147483647
func (e *GstKsVideoSrc) GetFps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The zero-based device index
// Default: -1, Min: -1, Max: 2147483647
func (e *GstKsVideoSrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

func (e *GstKsVideoSrc) GetDeviceIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The human-readable device name
// Default: NULL
func (e *GstKsVideoSrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

func (e *GstKsVideoSrc) GetDeviceName() (string, error) {
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

// The device path
// Default: NULL
func (e *GstKsVideoSrc) SetDevicePath(value string) error {
	return e.Element.SetProperty("device-path", value)
}

func (e *GstKsVideoSrc) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Enable logging of statistics
// Default: false
func (e *GstKsVideoSrc) SetDoStats(value bool) error {
	return e.Element.SetProperty("do-stats", value)
}

func (e *GstKsVideoSrc) GetDoStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable driver-specific quirks
// Default: true
func (e *GstKsVideoSrc) SetEnableQuirks(value bool) error {
	return e.Element.SetProperty("enable-quirks", value)
}

func (e *GstKsVideoSrc) GetEnableQuirks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-quirks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
