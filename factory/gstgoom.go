// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Takes frames of data and outputs video frames using the GOOM filter
type GstGoom struct {
	*GstAudioVisualizer
}

func NewGoom() (*GstGoom, error) {
	e, err := gst.NewElement("goom")
	if err != nil {
		return nil, err
	}
	return &GstGoom{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewGoomWithName(name string) (*GstGoom, error) {
	e, err := gst.NewElementWithName("goom", name)
	if err != nil {
		return nil, err
	}
	return &GstGoom{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}
