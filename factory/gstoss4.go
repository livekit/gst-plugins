// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output to a sound card via OSS version 4
type GstOss4Sink struct {
	*GstAudioSink
}

func NewOss4Sink() (*GstOss4Sink, error) {
	e, err := gst.NewElement("oss4sink")
	if err != nil {
		return nil, err
	}
	return &GstOss4Sink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewOss4SinkWithName(name string) (*GstOss4Sink, error) {
	e, err := gst.NewElementWithName("oss4sink", name)
	if err != nil {
		return nil, err
	}
	return &GstOss4Sink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Mute state of this stream
// Default: false
func (e *GstOss4Sink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstOss4Sink) GetMute() (bool, error) {
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

// Linear volume of this stream, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstOss4Sink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstOss4Sink) GetVolume() (float64, error) {
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

// OSS4 device (e.g. /dev/oss/hdaudio0/pcm0 or /dev/dspN) (NULL = use first available playback device)
// Default: NULL
func (e *GstOss4Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOss4Sink) GetDevice() (string, error) {
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
// Default: NULL
func (e *GstOss4Sink) GetDeviceName() (string, error) {
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

// Capture from a sound card via OSS version 4
type GstOss4Source struct {
	*GstAudioSrc
}

func NewOss4Source() (*GstOss4Source, error) {
	e, err := gst.NewElement("oss4src")
	if err != nil {
		return nil, err
	}
	return &GstOss4Source{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewOss4SourceWithName(name string) (*GstOss4Source, error) {
	e, err := gst.NewElementWithName("oss4src", name)
	if err != nil {
		return nil, err
	}
	return &GstOss4Source{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// OSS4 device (e.g. /dev/oss/hdaudio0/pcm0 or /dev/dspN) (NULL = use first available device)
// Default: NULL
func (e *GstOss4Source) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOss4Source) GetDevice() (string, error) {
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
// Default: NULL
func (e *GstOss4Source) GetDeviceName() (string, error) {
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
