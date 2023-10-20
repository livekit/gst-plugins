// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Output to a sound card in OS X
type GstOsxAudioSink struct {
	*GstAudioBaseSink
}

func NewOsxAudioSink() (*GstOsxAudioSink, error) {
	e, err := gst.NewElement("osxaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstOsxAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewOsxAudioSinkWithName(name string) (*GstOsxAudioSink, error) {
	e, err := gst.NewElementWithName("osxaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstOsxAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Volume of this stream
// Default: 1, Min: 0, Max: 1
func (e *GstOsxAudioSink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstOsxAudioSink) GetVolume() (float64, error) {
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

// Device ID of output device
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOsxAudioSink) SetDevice(value int) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOsxAudioSink) GetDevice() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Input from a sound card in OS X
type GstOsxAudioSrc struct {
	*GstAudioBaseSrc
}

func NewOsxAudioSrc() (*GstOsxAudioSrc, error) {
	e, err := gst.NewElement("osxaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstOsxAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

func NewOsxAudioSrcWithName(name string) (*GstOsxAudioSrc, error) {
	e, err := gst.NewElementWithName("osxaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstOsxAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}, nil
}

// Device ID of input device
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOsxAudioSrc) SetDevice(value int) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstOsxAudioSrc) GetDevice() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
