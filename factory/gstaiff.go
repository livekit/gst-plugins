// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Multiplex raw audio into AIFF
type GstAiffMux struct {
	*gst.Element
}

func NewAiffMux() (*GstAiffMux, error) {
	e, err := gst.NewElement("aiffmux")
	if err != nil {
		return nil, err
	}
	return &GstAiffMux{Element: e}, nil
}

func NewAiffMuxWithName(name string) (*GstAiffMux, error) {
	e, err := gst.NewElementWithName("aiffmux", name)
	if err != nil {
		return nil, err
	}
	return &GstAiffMux{Element: e}, nil
}

// Parse a .aiff file into raw audio
type GstAiffParse struct {
	*gst.Element
}

func NewAiffParse() (*GstAiffParse, error) {
	e, err := gst.NewElement("aiffparse")
	if err != nil {
		return nil, err
	}
	return &GstAiffParse{Element: e}, nil
}

func NewAiffParseWithName(name string) (*GstAiffParse, error) {
	e, err := gst.NewElementWithName("aiffparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAiffParse{Element: e}, nil
}
