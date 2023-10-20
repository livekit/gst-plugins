// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Simple stereo visualizer
type GstSpaceScope struct {
	*GstAudioVisualizer
}

func NewSpaceScope() (*GstSpaceScope, error) {
	e, err := gst.NewElement("spacescope")
	if err != nil {
		return nil, err
	}
	return &GstSpaceScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewSpaceScopeWithName(name string) (*GstSpaceScope, error) {
	e, err := gst.NewElementWithName("spacescope", name)
	if err != nil {
		return nil, err
	}
	return &GstSpaceScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

// Drawing styles for the space scope display.
// Default: dots (0)
func (e *GstSpaceScope) SetStyle(value GstSpaceScopeStyle) error {
	e.Element.SetArg("style", string(value))
	return nil
}

func (e *GstSpaceScope) GetStyle() (GstSpaceScopeStyle, error) {
	var value GstSpaceScopeStyle
	var ok bool
	v, err := e.Element.GetProperty("style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSpaceScopeStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSpaceScopeStyle")
	}
	return value, nil
}

// Simple frequency spectrum scope
type GstSpectraScope struct {
	*GstAudioVisualizer
}

func NewSpectraScope() (*GstSpectraScope, error) {
	e, err := gst.NewElement("spectrascope")
	if err != nil {
		return nil, err
	}
	return &GstSpectraScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewSpectraScopeWithName(name string) (*GstSpectraScope, error) {
	e, err := gst.NewElementWithName("spectrascope", name)
	if err != nil {
		return nil, err
	}
	return &GstSpectraScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

// Creates video visualizations of audio input, using stereo and pitch information
type GstSynaeScope struct {
	*GstAudioVisualizer
}

func NewSynaeScope() (*GstSynaeScope, error) {
	e, err := gst.NewElement("synaescope")
	if err != nil {
		return nil, err
	}
	return &GstSynaeScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewSynaeScopeWithName(name string) (*GstSynaeScope, error) {
	e, err := gst.NewElementWithName("synaescope", name)
	if err != nil {
		return nil, err
	}
	return &GstSynaeScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

// Simple waveform oscilloscope
type GstWaveScope struct {
	*GstAudioVisualizer
}

func NewWaveScope() (*GstWaveScope, error) {
	e, err := gst.NewElement("wavescope")
	if err != nil {
		return nil, err
	}
	return &GstWaveScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewWaveScopeWithName(name string) (*GstWaveScope, error) {
	e, err := gst.NewElementWithName("wavescope", name)
	if err != nil {
		return nil, err
	}
	return &GstWaveScope{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

// Drawing styles for the wave form display.
// Default: dots (0)
func (e *GstWaveScope) SetStyle(value GstWaveScopeStyle) error {
	e.Element.SetArg("style", string(value))
	return nil
}

func (e *GstWaveScope) GetStyle() (GstWaveScopeStyle, error) {
	var value GstWaveScopeStyle
	var ok bool
	v, err := e.Element.GetProperty("style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWaveScopeStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWaveScopeStyle")
	}
	return value, nil
}

type GstSpaceScopeStyle string

const (
	GstSpaceScopeStyleDots       GstSpaceScopeStyle = "dots"        // draw dots (default)
	GstSpaceScopeStyleLines      GstSpaceScopeStyle = "lines"       // draw lines
	GstSpaceScopeStyleColorDots  GstSpaceScopeStyle = "color-dots"  // draw color dots
	GstSpaceScopeStyleColorLines GstSpaceScopeStyle = "color-lines" // draw color lines
)

type GstWaveScopeStyle string

const (
	GstWaveScopeStyleDots       GstWaveScopeStyle = "dots"        // draw dots (default)
	GstWaveScopeStyleLines      GstWaveScopeStyle = "lines"       // draw lines
	GstWaveScopeStyleColorDots  GstWaveScopeStyle = "color-dots"  // draw color dots
	GstWaveScopeStyleColorLines GstWaveScopeStyle = "color-lines" // draw color lines
)
