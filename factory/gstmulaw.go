// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Convert 16bit PCM to 8bit mu law
type GstMuLawEnc struct {
	*GstAudioEncoder
}

func NewMuLawEnc() (*GstMuLawEnc, error) {
	e, err := gst.NewElement("mulawenc")
	if err != nil {
		return nil, err
	}
	return &GstMuLawEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewMuLawEncWithName(name string) (*GstMuLawEnc, error) {
	e, err := gst.NewElementWithName("mulawenc", name)
	if err != nil {
		return nil, err
	}
	return &GstMuLawEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Convert 8bit mu law to 16bit PCM
type GstMuLawDec struct {
	*GstAudioDecoder
}

func NewMuLawDec() (*GstMuLawDec, error) {
	e, err := gst.NewElement("mulawdec")
	if err != nil {
		return nil, err
	}
	return &GstMuLawDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewMuLawDecWithName(name string) (*GstMuLawDec, error) {
	e, err := gst.NewElementWithName("mulawdec", name)
	if err != nil {
		return nil, err
	}
	return &GstMuLawDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
