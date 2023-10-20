// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Adds an APEv2 header to the beginning of files using taglib
type GstApev2Mux struct {
	*GstTagMux
}

func NewApev2Mux() (*GstApev2Mux, error) {
	e, err := gst.NewElement("apev2mux")
	if err != nil {
		return nil, err
	}
	return &GstApev2Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}

func NewApev2MuxWithName(name string) (*GstApev2Mux, error) {
	e, err := gst.NewElementWithName("apev2mux", name)
	if err != nil {
		return nil, err
	}
	return &GstApev2Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}

// Adds an ID3v2 header to the beginning of MP3 files using taglib
type GstId3v2Mux struct {
	*GstTagMux
}

func NewId3v2Mux() (*GstId3v2Mux, error) {
	e, err := gst.NewElement("id3v2mux")
	if err != nil {
		return nil, err
	}
	return &GstId3v2Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}

func NewId3v2MuxWithName(name string) (*GstId3v2Mux, error) {
	e, err := gst.NewElementWithName("id3v2mux", name)
	if err != nil {
		return nil, err
	}
	return &GstId3v2Mux{GstTagMux: &GstTagMux{Element: e}}, nil
}
