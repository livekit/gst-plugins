// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Source for DV video data from firewire port
type GstDV1394Src struct {
	*base.GstPushSrc
}

func NewDV1394Src() (*GstDV1394Src, error) {
	e, err := gst.NewElement("dv1394src")
	if err != nil {
		return nil, err
	}
	return &GstDV1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDV1394SrcWithName(name string) (*GstDV1394Src, error) {
	e, err := gst.NewElementWithName("dv1394src", name)
	if err != nil {
		return nil, err
	}
	return &GstDV1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// send n consecutive frames after skipping
// Default: 1, Min: 1, Max: 2147483647
func (e *GstDV1394Src) SetConsecutive(value int) error {
	return e.Element.SetProperty("consecutive", value)
}

func (e *GstDV1394Src) GetConsecutive() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("consecutive")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// user-friendly name of the device
// Default: Default
func (e *GstDV1394Src) GetDeviceName() (string, error) {
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

// drop incomplete frames
// Default: true
func (e *GstDV1394Src) SetDropIncomplete(value bool) error {
	return e.Element.SetProperty("drop-incomplete", value)
}

func (e *GstDV1394Src) GetDropIncomplete() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-incomplete")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// select one of multiple DV devices by its GUID. use a hexadecimal like 0xhhhhhhhhhhhhhhhh. (0 = no guid)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDV1394Src) SetGuid(value uint64) error {
	return e.Element.SetProperty("guid", value)
}

func (e *GstDV1394Src) GetGuid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("guid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Port number (-1 automatic)
// Default: -1, Min: -1, Max: 16
func (e *GstDV1394Src) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstDV1394Src) GetPort() (int, error) {
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

// skip n frames
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDV1394Src) SetSkip(value int) error {
	return e.Element.SetProperty("skip", value)
}

func (e *GstDV1394Src) GetSkip() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Use AV/C VTR control
// Default: true
func (e *GstDV1394Src) SetUseAvc(value bool) error {
	return e.Element.SetProperty("use-avc", value)
}

func (e *GstDV1394Src) GetUseAvc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-avc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Channel number for listening
// Default: 63, Min: 0, Max: 64
func (e *GstDV1394Src) SetChannel(value int) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstDV1394Src) GetChannel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Source for MPEG-TS video data from firewire port
type GstHDV1394Src struct {
	*base.GstPushSrc
}

func NewHDV1394Src() (*GstHDV1394Src, error) {
	e, err := gst.NewElement("hdv1394src")
	if err != nil {
		return nil, err
	}
	return &GstHDV1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewHDV1394SrcWithName(name string) (*GstHDV1394Src, error) {
	e, err := gst.NewElementWithName("hdv1394src", name)
	if err != nil {
		return nil, err
	}
	return &GstHDV1394Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Channel number for listening
// Default: 63, Min: 0, Max: 64
func (e *GstHDV1394Src) SetChannel(value int) error {
	return e.Element.SetProperty("channel", value)
}

func (e *GstHDV1394Src) GetChannel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// user-friendly name of the device
// Default: Default
func (e *GstHDV1394Src) GetDeviceName() (string, error) {
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

// select one of multiple DV devices by its GUID. use a hexadecimal like 0xhhhhhhhhhhhhhhhh. (0 = no guid)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstHDV1394Src) SetGuid(value uint64) error {
	return e.Element.SetProperty("guid", value)
}

func (e *GstHDV1394Src) GetGuid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("guid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Port number (-1 automatic)
// Default: -1, Min: -1, Max: 16
func (e *GstHDV1394Src) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstHDV1394Src) GetPort() (int, error) {
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

// Use AV/C VTR control
// Default: true
func (e *GstHDV1394Src) SetUseAvc(value bool) error {
	return e.Element.SetProperty("use-avc", value)
}

func (e *GstHDV1394Src) GetUseAvc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-avc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
