// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Midi Parser Element
type GstMidiParse struct {
	*gst.Element
}

func NewMidiParse() (*GstMidiParse, error) {
	e, err := gst.NewElement("midiparse")
	if err != nil {
		return nil, err
	}
	return &GstMidiParse{Element: e}, nil
}

func NewMidiParseWithName(name string) (*GstMidiParse, error) {
	e, err := gst.NewElementWithName("midiparse", name)
	if err != nil {
		return nil, err
	}
	return &GstMidiParse{Element: e}, nil
}
