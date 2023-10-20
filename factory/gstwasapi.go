// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Stream audio to an audio capture device through WASAPI
type GstWasapiSink struct {
	*GstAudioSink
}

func NewWasapiSink() (*GstWasapiSink, error) {
	e, err := gst.NewElement("wasapisink")
	if err != nil {
		return nil, err
	}
	return &GstWasapiSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewWasapiSinkWithName(name string) (*GstWasapiSink, error) {
	e, err := gst.NewElementWithName("wasapisink", name)
	if err != nil {
		return nil, err
	}
	return &GstWasapiSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// WASAPI device endpoint ID as provided by IMMDevice::GetId
// Default: NULL
func (e *GstWasapiSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstWasapiSink) GetDevice() (string, error) {
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

// Open the device in exclusive mode
// Default: false
func (e *GstWasapiSink) SetExclusive(value bool) error {
	return e.Element.SetProperty("exclusive", value)
}

func (e *GstWasapiSink) GetExclusive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("exclusive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Optimize all settings for lowest latency. Always safe to enable.
// Default: false
func (e *GstWasapiSink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstWasapiSink) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Mute state of this stream
// Default: false
func (e *GstWasapiSink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstWasapiSink) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Role of the device: communications, multimedia, etc
// Default: console (0)
func (e *GstWasapiSink) SetRole(value GstWasapiDeviceRole) error {
	e.Element.SetArg("role", string(value))
	return nil
}

func (e *GstWasapiSink) GetRole() (GstWasapiDeviceRole, error) {
	var value GstWasapiDeviceRole
	var ok bool
	v, err := e.Element.GetProperty("role")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWasapiDeviceRole)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWasapiDeviceRole")
	}
	return value, nil
}

// Use the Windows 10 AudioClient3 API when available and if the low-latency property is set to TRUE
// Default: true
func (e *GstWasapiSink) SetUseAudioclient3(value bool) error {
	return e.Element.SetProperty("use-audioclient3", value)
}

func (e *GstWasapiSink) GetUseAudioclient3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-audioclient3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Stream audio from an audio capture device through WASAPI
type GstWasapiSrc struct {
	*GstAudioSrc
}

func NewWasapiSrc() (*GstWasapiSrc, error) {
	e, err := gst.NewElement("wasapisrc")
	if err != nil {
		return nil, err
	}
	return &GstWasapiSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewWasapiSrcWithName(name string) (*GstWasapiSrc, error) {
	e, err := gst.NewElementWithName("wasapisrc", name)
	if err != nil {
		return nil, err
	}
	return &GstWasapiSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// Open the sink device for loopback recording
// Default: false
func (e *GstWasapiSrc) SetLoopback(value bool) error {
	return e.Element.SetProperty("loopback", value)
}

func (e *GstWasapiSrc) GetLoopback() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loopback")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Optimize all settings for lowest latency. Always safe to enable.
// Default: false
func (e *GstWasapiSrc) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstWasapiSrc) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Role of the device: communications, multimedia, etc
// Default: console (0)
func (e *GstWasapiSrc) SetRole(value GstWasapiDeviceRole) error {
	e.Element.SetArg("role", string(value))
	return nil
}

func (e *GstWasapiSrc) GetRole() (GstWasapiDeviceRole, error) {
	var value GstWasapiDeviceRole
	var ok bool
	v, err := e.Element.GetProperty("role")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWasapiDeviceRole)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWasapiDeviceRole")
	}
	return value, nil
}

// Whether to use the Windows 10 AudioClient3 API when available
// Default: false
func (e *GstWasapiSrc) SetUseAudioclient3(value bool) error {
	return e.Element.SetProperty("use-audioclient3", value)
}

func (e *GstWasapiSrc) GetUseAudioclient3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-audioclient3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// WASAPI device endpoint ID as provided by IMMDevice::GetId
// Default: NULL
func (e *GstWasapiSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstWasapiSrc) GetDevice() (string, error) {
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

// Open the device in exclusive mode
// Default: false
func (e *GstWasapiSrc) SetExclusive(value bool) error {
	return e.Element.SetProperty("exclusive", value)
}

func (e *GstWasapiSrc) GetExclusive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("exclusive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstWasapiDeviceRole string

const (
	GstWasapiDeviceRoleConsole    GstWasapiDeviceRole = "console"    // Games, system notifications, voice commands
	GstWasapiDeviceRoleMultimedia GstWasapiDeviceRole = "multimedia" // Music, movies, recorded media
	GstWasapiDeviceRoleComms      GstWasapiDeviceRole = "comms"      // Voice communications
)
