// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Free MPEG-2/4 AAC decoder
type GstFaad struct {
	*GstAudioDecoder
}

func NewFaad() (*GstFaad, error) {
	e, err := gst.NewElement("faad")
	if err != nil {
		return nil, err
	}
	return &GstFaad{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewFaadWithName(name string) (*GstFaad, error) {
	e, err := gst.NewElementWithName("faad", name)
	if err != nil {
		return nil, err
	}
	return &GstFaad{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
