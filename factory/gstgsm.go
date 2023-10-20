// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decodes GSM encoded audio
type GstGSMDec struct {
	*GstAudioDecoder
}

func NewGSMDec() (*GstGSMDec, error) {
	e, err := gst.NewElement("gsmdec")
	if err != nil {
		return nil, err
	}
	return &GstGSMDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewGSMDecWithName(name string) (*GstGSMDec, error) {
	e, err := gst.NewElementWithName("gsmdec", name)
	if err != nil {
		return nil, err
	}
	return &GstGSMDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encodes GSM audio
type GstGSMEnc struct {
	*GstAudioEncoder
}

func NewGSMEnc() (*GstGSMEnc, error) {
	e, err := gst.NewElement("gsmenc")
	if err != nil {
		return nil, err
	}
	return &GstGSMEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewGSMEncWithName(name string) (*GstGSMEnc, error) {
	e, err := gst.NewElementWithName("gsmenc", name)
	if err != nil {
		return nil, err
	}
	return &GstGSMEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}
