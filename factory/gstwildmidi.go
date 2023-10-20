// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes MIDI music using WildMidi
type GstWildmidiDec struct {
	*GstNonstreamAudioDecoder
}

func NewWildmidiDec() (*GstWildmidiDec, error) {
	e, err := gst.NewElement("wildmididec")
	if err != nil {
		return nil, err
	}
	return &GstWildmidiDec{GstNonstreamAudioDecoder: &GstNonstreamAudioDecoder{Element: e}}, nil
}

func NewWildmidiDecWithName(name string) (*GstWildmidiDec, error) {
	e, err := gst.NewElementWithName("wildmididec", name)
	if err != nil {
		return nil, err
	}
	return &GstWildmidiDec{GstNonstreamAudioDecoder: &GstNonstreamAudioDecoder{Element: e}}, nil
}

// Use a logarithmic volume scale if set to TRUE, or a linear scale if set to FALSE
// Default: true
func (e *GstWildmidiDec) SetLogVolumeScale(value bool) error {
	return e.Element.SetProperty("log-volume-scale", value)
}

func (e *GstWildmidiDec) GetLogVolumeScale() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("log-volume-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Size of each output buffer, in samples (actual size can be smaller than this during flush or EOS)
// Default: 1024, Min: 1, Max: 1073741823
func (e *GstWildmidiDec) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

func (e *GstWildmidiDec) GetOutputBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("output-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether or not to enable the WildMidi 8 reflection reverb engine to add more depth to the sound
// Default: false
func (e *GstWildmidiDec) SetReverb(value bool) error {
	return e.Element.SetProperty("reverb", value)
}

func (e *GstWildmidiDec) GetReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use enhanced resampling if set to TRUE, or linear interpolation if set to FALSE
// Default: true
func (e *GstWildmidiDec) SetEnhancedResampling(value bool) error {
	return e.Element.SetProperty("enhanced-resampling", value)
}

func (e *GstWildmidiDec) GetEnhancedResampling() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enhanced-resampling")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
