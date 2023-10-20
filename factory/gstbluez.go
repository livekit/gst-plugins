// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Plays audio to an A2DP device
type GstAvdtpSink struct {
	*base.GstBaseSink
}

func NewAvdtpSink() (*GstAvdtpSink, error) {
	e, err := gst.NewElement("avdtpsink")
	if err != nil {
		return nil, err
	}
	return &GstAvdtpSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAvdtpSinkWithName(name string) (*GstAvdtpSink, error) {
	e, err := gst.NewElementWithName("avdtpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstAvdtpSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Automatically attempt to connect to device
// Default: true
func (e *GstAvdtpSink) SetAutoConnect(value bool) error {
	return e.Element.SetProperty("auto-connect", value)
}

func (e *GstAvdtpSink) GetAutoConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bluetooth remote device address
// Default: NULL
func (e *GstAvdtpSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstAvdtpSink) GetDevice() (string, error) {
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

// Use configured transport
// Default: NULL
func (e *GstAvdtpSink) SetTransport(value string) error {
	return e.Element.SetProperty("transport", value)
}

func (e *GstAvdtpSink) GetTransport() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Receives audio from an A2DP device
type GstAvdtpSrc struct {
	*base.GstBaseSrc
}

func NewAvdtpSrc() (*GstAvdtpSrc, error) {
	e, err := gst.NewElement("avdtpsrc")
	if err != nil {
		return nil, err
	}
	return &GstAvdtpSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewAvdtpSrcWithName(name string) (*GstAvdtpSrc, error) {
	e, err := gst.NewElementWithName("avdtpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAvdtpSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Use configured transport
// Default: NULL
func (e *GstAvdtpSrc) SetTransport(value string) error {
	return e.Element.SetProperty("transport", value)
}

func (e *GstAvdtpSrc) GetTransport() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Volume of the transport (only valid if transport is acquired)
// Default: 127, Min: 0, Max: 127
func (e *GstAvdtpSrc) SetTransportVolume(value uint) error {
	return e.Element.SetProperty("transport-volume", value)
}

func (e *GstAvdtpSrc) GetTransportVolume() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("transport-volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Plays audio to an A2DP device
type GstA2dpSink struct {
	*gst.Bin
}

func NewA2dpSink() (*GstA2dpSink, error) {
	e, err := gst.NewElement("a2dpsink")
	if err != nil {
		return nil, err
	}
	return &GstA2dpSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewA2dpSinkWithName(name string) (*GstA2dpSink, error) {
	e, err := gst.NewElementWithName("a2dpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstA2dpSink{Bin: &gst.Bin{Element: e}}, nil
}

// Automatically attempt to connect to device
// Default: true
func (e *GstA2dpSink) SetAutoConnect(value bool) error {
	return e.Element.SetProperty("auto-connect", value)
}

func (e *GstA2dpSink) GetAutoConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Bluetooth remote device address
// Default: NULL
func (e *GstA2dpSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstA2dpSink) GetDevice() (string, error) {
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

// Use configured transport
// Default: NULL
func (e *GstA2dpSink) SetTransport(value string) error {
	return e.Element.SetProperty("transport", value)
}

func (e *GstA2dpSink) GetTransport() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("transport")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
