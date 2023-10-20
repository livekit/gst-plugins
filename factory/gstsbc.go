// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decode an SBC audio stream
type GstSbcDec struct {
	*GstAudioDecoder
}

func NewSbcDec() (*GstSbcDec, error) {
	e, err := gst.NewElement("sbcdec")
	if err != nil {
		return nil, err
	}
	return &GstSbcDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewSbcDecWithName(name string) (*GstSbcDec, error) {
	e, err := gst.NewElementWithName("sbcdec", name)
	if err != nil {
		return nil, err
	}
	return &GstSbcDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encode an SBC audio stream
type GstSbcEnc struct {
	*GstAudioEncoder
}

func NewSbcEnc() (*GstSbcEnc, error) {
	e, err := gst.NewElement("sbcenc")
	if err != nil {
		return nil, err
	}
	return &GstSbcEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewSbcEncWithName(name string) (*GstSbcEnc, error) {
	e, err := gst.NewElementWithName("sbcenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSbcEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}
