// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decodes mp3 streams using the mpg123 library
type GstMpg123AudioDec struct {
	*GstAudioDecoder
}

func NewMpg123AudioDec() (*GstMpg123AudioDec, error) {
	e, err := gst.NewElement("mpg123audiodec")
	if err != nil {
		return nil, err
	}
	return &GstMpg123AudioDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewMpg123AudioDecWithName(name string) (*GstMpg123AudioDec, error) {
	e, err := gst.NewElementWithName("mpg123audiodec", name)
	if err != nil {
		return nil, err
	}
	return &GstMpg123AudioDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
