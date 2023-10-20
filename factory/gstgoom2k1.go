// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Takes frames of data and outputs video frames using the GOOM 2k1 filter
type GstGoom2k1 struct {
	*GstAudioVisualizer
}

func NewGoom2k1() (*GstGoom2k1, error) {
	e, err := gst.NewElement("goom2k1")
	if err != nil {
		return nil, err
	}
	return &GstGoom2k1{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}

func NewGoom2k1WithName(name string) (*GstGoom2k1, error) {
	e, err := gst.NewElementWithName("goom2k1", name)
	if err != nil {
		return nil, err
	}
	return &GstGoom2k1{GstAudioVisualizer: &GstAudioVisualizer{Element: e}}, nil
}
