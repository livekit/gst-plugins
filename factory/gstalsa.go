// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Read from a sound card via ALSA
type GstAlsaSrc struct {
	*GstAudioSrc
}

func NewAlsaSrc() (*GstAlsaSrc, error) {
	e, err := gst.NewElement("alsasrc")
	if err != nil {
		return nil, err
	}
	return &GstAlsaSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewAlsaSrcWithName(name string) (*GstAlsaSrc, error) {
	e, err := gst.NewElementWithName("alsasrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAlsaSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// Human-readable name of the sound card

func (e *GstAlsaSrc) GetCardName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("card-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// ALSA device, as defined in an asound configuration file
// Default: default
func (e *GstAlsaSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstAlsaSrc) GetDevice() (string, error) {
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
func (e *GstAlsaSrc) GetDeviceName() (string, error) {
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

// Use driver timestamps or the pipeline clock timestamps
// Default: true
func (e *GstAlsaSrc) SetUseDriverTimestamps(value bool) error {
	return e.Element.SetProperty("use-driver-timestamps", value)
}

func (e *GstAlsaSrc) GetUseDriverTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-driver-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Push ALSA MIDI sequencer events around
type GstAlsaMidiSrc struct {
	*base.GstPushSrc
}

func NewAlsaMidiSrc() (*GstAlsaMidiSrc, error) {
	e, err := gst.NewElement("alsamidisrc")
	if err != nil {
		return nil, err
	}
	return &GstAlsaMidiSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewAlsaMidiSrcWithName(name string) (*GstAlsaMidiSrc, error) {
	e, err := gst.NewElementWithName("alsamidisrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAlsaMidiSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Comma separated list of sequencer ports (e.g. client:port,...)
// Default: NULL
func (e *GstAlsaMidiSrc) SetPorts(value string) error {
	return e.Element.SetProperty("ports", value)
}

func (e *GstAlsaMidiSrc) GetPorts() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ports")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Output to a sound card via ALSA
type GstAlsaSink struct {
	*GstAudioSink
}

func NewAlsaSink() (*GstAlsaSink, error) {
	e, err := gst.NewElement("alsasink")
	if err != nil {
		return nil, err
	}
	return &GstAlsaSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewAlsaSinkWithName(name string) (*GstAlsaSink, error) {
	e, err := gst.NewElementWithName("alsasink", name)
	if err != nil {
		return nil, err
	}
	return &GstAlsaSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Human-readable name of the sound card

func (e *GstAlsaSink) GetCardName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("card-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// ALSA device, as defined in an asound configuration file
// Default: default
func (e *GstAlsaSink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstAlsaSink) GetDevice() (string, error) {
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
func (e *GstAlsaSink) GetDeviceName() (string, error) {
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
