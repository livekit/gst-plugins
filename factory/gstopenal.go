// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output audio through OpenAL
type GstOpenALSink struct {
	*GstAudioSink
}

func NewOpenALSink() (*GstOpenALSink, error) {
	e, err := gst.NewElement("openalsink")
	if err != nil {
		return nil, err
	}
	return &GstOpenALSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewOpenALSinkWithName(name string) (*GstOpenALSink, error) {
	e, err := gst.NewElementWithName("openalsink", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenALSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Human-readable name of the device
// Default: NULL
func (e *GstOpenALSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOpenALSink) GetDevice() (string, error) {
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

// Human-readable name of the opened device

func (e *GstOpenALSink) GetDeviceName() (string, error) {
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

// User context

func (e *GstOpenALSink) SetUserContext(value interface{}) error {
	return e.Element.SetProperty("user-context", value)
}

func (e *GstOpenALSink) GetUserContext() (interface{}, error) {
	return e.Element.GetProperty("user-context")
}

// User device

func (e *GstOpenALSink) SetUserDevice(value interface{}) error {
	return e.Element.SetProperty("user-device", value)
}

func (e *GstOpenALSink) GetUserDevice() (interface{}, error) {
	return e.Element.GetProperty("user-device")
}

// User source
// Default: 0, Min: 0, Max: -1
func (e *GstOpenALSink) SetUserSource(value uint) error {
	return e.Element.SetProperty("user-source", value)
}

func (e *GstOpenALSink) GetUserSource() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("user-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Input audio through OpenAL
type GstOpenalSrc struct {
	*GstAudioSrc
}

func NewOpenalSrc() (*GstOpenalSrc, error) {
	e, err := gst.NewElement("openalsrc")
	if err != nil {
		return nil, err
	}
	return &GstOpenalSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewOpenalSrcWithName(name string) (*GstOpenalSrc, error) {
	e, err := gst.NewElementWithName("openalsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenalSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// User device, default device if NULL
// Default: NULL
func (e *GstOpenalSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOpenalSrc) GetDevice() (string, error) {
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

// Human-readable name of the device
// Default: NULL
func (e *GstOpenalSrc) GetDeviceName() (string, error) {
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
