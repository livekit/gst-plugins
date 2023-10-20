// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Stream audio to an audio capture device through WASAPI
type GstWasapi2Sink struct {
	*GstAudioBaseSink
}

func NewWasapi2Sink() (*GstWasapi2Sink, error) {
	e, err := gst.NewElement("wasapi2sink")
	if err != nil {
		return nil, err
	}
	return &GstWasapi2Sink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewWasapi2SinkWithName(name string) (*GstWasapi2Sink, error) {
	e, err := gst.NewElementWithName("wasapi2sink", name)
	if err != nil {
		return nil, err
	}
	return &GstWasapi2Sink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Optimize all settings for lowest latency. Always safe to enable.
// Default: false
func (e *GstWasapi2Sink) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstWasapi2Sink) GetLowLatency() (bool, error) {
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
func (e *GstWasapi2Sink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstWasapi2Sink) GetMute() (bool, error) {
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

// Volume of this stream
// Default: 1, Min: 0, Max: 1
func (e *GstWasapi2Sink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstWasapi2Sink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Audio device ID as provided by Windows.Devices.Enumeration.DeviceInformation.Id
// Default: NULL
func (e *GstWasapi2Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstWasapi2Sink) GetDevice() (string, error) {
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

// ICoreDispatcher COM object to use. In order for application to ask permission of audio device, device activation should be running on UI thread via ICoreDispatcher. This element will increase the reference count of given ICoreDispatcher and release it after use. Therefore, caller does not need to consider additional reference count management

func (e *GstWasapi2Sink) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// Stream audio from an audio capture device through WASAPI
type GstWasapi2Src struct {
	*GstAudioBaseSrc
}

func NewWasapi2Src() (*GstWasapi2Src, error) {
	e, err := gst.NewElement("wasapi2src")
	if err != nil {
		return nil, err
	}
	return &GstWasapi2Src{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewWasapi2SrcWithName(name string) (*GstWasapi2Src, error) {
	e, err := gst.NewElementWithName("wasapi2src", name)
	if err != nil {
		return nil, err
	}
	return &GstWasapi2Src{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Loopback mode to use
// Default: default (0)
func (e *GstWasapi2Src) SetLoopbackMode(value GstWasapi2SrcLoopbackMode) error {
	e.Element.SetArg("loopback-mode", string(value))
	return nil
}

func (e *GstWasapi2Src) GetLoopbackMode() (GstWasapi2SrcLoopbackMode, error) {
	var value GstWasapi2SrcLoopbackMode
	var ok bool
	v, err := e.Element.GetProperty("loopback-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWasapi2SrcLoopbackMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWasapi2SrcLoopbackMode")
	}
	return value, nil
}

// When loopback recording, if the device is muted, inject silence in the pipeline
// Default: false
func (e *GstWasapi2Src) SetLoopbackSilenceOnDeviceMute(value bool) error {
	return e.Element.SetProperty("loopback-silence-on-device-mute", value)
}

func (e *GstWasapi2Src) GetLoopbackSilenceOnDeviceMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loopback-silence-on-device-mute")
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
func (e *GstWasapi2Src) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstWasapi2Src) GetLowLatency() (bool, error) {
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
func (e *GstWasapi2Src) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstWasapi2Src) GetMute() (bool, error) {
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

// Volume of this stream
// Default: 1, Min: 0, Max: 1
func (e *GstWasapi2Src) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstWasapi2Src) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Audio device ID as provided by Windows.Devices.Enumeration.DeviceInformation.Id
// Default: NULL
func (e *GstWasapi2Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstWasapi2Src) GetDevice() (string, error) {
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

// ICoreDispatcher COM object to use. In order for application to ask permission of audio device, device activation should be running on UI thread via ICoreDispatcher. This element will increase the reference count of given ICoreDispatcher and release it after use. Therefore, caller does not need to consider additional reference count management

func (e *GstWasapi2Src) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// Open render device for loopback recording
// Default: false
func (e *GstWasapi2Src) SetLoopback(value bool) error {
	return e.Element.SetProperty("loopback", value)
}

func (e *GstWasapi2Src) GetLoopback() (bool, error) {
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

// Process ID to be recorded or excluded for process loopback mode
// Default: 0, Min: 0, Max: -1
func (e *GstWasapi2Src) SetLoopbackTargetPid(value uint) error {
	return e.Element.SetProperty("loopback-target-pid", value)
}

func (e *GstWasapi2Src) GetLoopbackTargetPid() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("loopback-target-pid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstWasapi2SrcLoopbackMode string

const (
	GstWasapi2SrcLoopbackModeDefault            GstWasapi2SrcLoopbackMode = "default"              // Default
	GstWasapi2SrcLoopbackModeIncludeProcessTree GstWasapi2SrcLoopbackMode = "include-process-tree" // Include process and its child processes
	GstWasapi2SrcLoopbackModeExcludeProcessTree GstWasapi2SrcLoopbackMode = "exclude-process-tree" // Exclude process and its child processes
)
