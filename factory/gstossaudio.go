// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output to a sound card via OSS
type GstOssSink struct {
	*GstAudioSink
}

func NewOssSink() (*GstOssSink, error) {
	e, err := gst.NewElement("osssink")
	if err != nil {
		return nil, err
	}
	return &GstOssSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewOssSinkWithName(name string) (*GstOssSink, error) {
	e, err := gst.NewElementWithName("osssink", name)
	if err != nil {
		return nil, err
	}
	return &GstOssSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// OSS device (usually /dev/dspN)
// Default: /dev/dsp
func (e *GstOssSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOssSink) GetDevice() (string, error) {
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

// Capture from a sound card via OSS
type GstOssSrc struct {
	*GstAudioSrc
}

func NewOssSrc() (*GstOssSrc, error) {
	e, err := gst.NewElement("osssrc")
	if err != nil {
		return nil, err
	}
	return &GstOssSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewOssSrcWithName(name string) (*GstOssSrc, error) {
	e, err := gst.NewElementWithName("osssrc", name)
	if err != nil {
		return nil, err
	}
	return &GstOssSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// OSS device (usually /dev/dspN)
// Default: /dev/dsp
func (e *GstOssSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOssSrc) GetDevice() (string, error) {
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

// Human-readable name of the sound device

func (e *GstOssSrc) GetDeviceName() (string, error) {
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
