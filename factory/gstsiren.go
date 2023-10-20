// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode streams encoded with the Siren7 codec into 16bit PCM
type GstSirenDec struct {
	*GstAudioDecoder
}

func NewSirenDec() (*GstSirenDec, error) {
	e, err := gst.NewElement("sirendec")
	if err != nil {
		return nil, err
	}
	return &GstSirenDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewSirenDecWithName(name string) (*GstSirenDec, error) {
	e, err := gst.NewElementWithName("sirendec", name)
	if err != nil {
		return nil, err
	}
	return &GstSirenDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encode 16bit PCM streams into the Siren7 codec
type GstSirenEnc struct {
	*GstAudioEncoder
}

func NewSirenEnc() (*GstSirenEnc, error) {
	e, err := gst.NewElement("sirenenc")
	if err != nil {
		return nil, err
	}
	return &GstSirenEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewSirenEncWithName(name string) (*GstSirenEnc, error) {
	e, err := gst.NewElementWithName("sirenenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSirenEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}
