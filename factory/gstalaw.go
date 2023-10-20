// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Convert 8bit A law to 16bit PCM
type GstALawDec struct {
	*GstAudioDecoder
}

func NewALawDec() (*GstALawDec, error) {
	e, err := gst.NewElement("alawdec")
	if err != nil {
		return nil, err
	}
	return &GstALawDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewALawDecWithName(name string) (*GstALawDec, error) {
	e, err := gst.NewElementWithName("alawdec", name)
	if err != nil {
		return nil, err
	}
	return &GstALawDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Convert 16bit PCM to 8bit A law
type GstALawEnc struct {
	*GstAudioEncoder
}

func NewALawEnc() (*GstALawEnc, error) {
	e, err := gst.NewElement("alawenc")
	if err != nil {
		return nil, err
	}
	return &GstALawEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewALawEncWithName(name string) (*GstALawEnc, error) {
	e, err := gst.NewElementWithName("alawenc", name)
	if err != nil {
		return nil, err
	}
	return &GstALawEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}
