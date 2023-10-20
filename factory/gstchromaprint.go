// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Find an audio fingerprint using the Chromaprint library
type GstChromaprint struct {
	*GstAudioFilter
}

func NewChromaprint() (*GstChromaprint, error) {
	e, err := gst.NewElement("chromaprint")
	if err != nil {
		return nil, err
	}
	return &GstChromaprint{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewChromaprintWithName(name string) (*GstChromaprint, error) {
	e, err := gst.NewElementWithName("chromaprint", name)
	if err != nil {
		return nil, err
	}
	return &GstChromaprint{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Number of seconds of audio to use for fingerprinting
// Default: 120, Min: 0, Max: -1
func (e *GstChromaprint) SetDuration(value uint) error {
	return e.Element.SetProperty("duration", value)
}

func (e *GstChromaprint) GetDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Resulting fingerprint
// Default: NULL
func (e *GstChromaprint) GetFingerprint() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fingerprint")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
