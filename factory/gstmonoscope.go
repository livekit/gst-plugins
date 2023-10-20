// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Displays a highly stabilised waveform of audio input
type GstMonoscope struct {
	*gst.Element
}

func NewMonoscope() (*GstMonoscope, error) {
	e, err := gst.NewElement("monoscope")
	if err != nil {
		return nil, err
	}
	return &GstMonoscope{Element: e}, nil
}

func NewMonoscopeWithName(name string) (*GstMonoscope, error) {
	e, err := gst.NewElementWithName("monoscope", name)
	if err != nil {
		return nil, err
	}
	return &GstMonoscope{Element: e}, nil
}
