// source: gst-plugins-ugly

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode DVD LPCM frames into standard PCM audio
type GstDvdLpcmDec struct {
	*GstAudioDecoder
}

func NewDvdLpcmDec() (*GstDvdLpcmDec, error) {
	e, err := gst.NewElement("dvdlpcmdec")
	if err != nil {
		return nil, err
	}
	return &GstDvdLpcmDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewDvdLpcmDecWithName(name string) (*GstDvdLpcmDec, error) {
	e, err := gst.NewElementWithName("dvdlpcmdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDvdLpcmDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
