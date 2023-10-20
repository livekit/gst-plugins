// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode MS and IMA ADPCM audio
type ADPCMDec struct {
	*GstAudioDecoder
}

func NewADPCMDec() (*ADPCMDec, error) {
	e, err := gst.NewElement("adpcmdec")
	if err != nil {
		return nil, err
	}
	return &ADPCMDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewADPCMDecWithName(name string) (*ADPCMDec, error) {
	e, err := gst.NewElementWithName("adpcmdec", name)
	if err != nil {
		return nil, err
	}
	return &ADPCMDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
