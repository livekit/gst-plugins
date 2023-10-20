// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decodes an LC3 Audio stream to raw audio
type GstLc3Dec struct {
	*GstAudioDecoder
}

func NewLc3Dec() (*GstLc3Dec, error) {
	e, err := gst.NewElement("lc3dec")
	if err != nil {
		return nil, err
	}
	return &GstLc3Dec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewLc3DecWithName(name string) (*GstLc3Dec, error) {
	e, err := gst.NewElementWithName("lc3dec", name)
	if err != nil {
		return nil, err
	}
	return &GstLc3Dec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Encodes a raw audio stream to LC3
type GstLc3Enc struct {
	*GstAudioEncoder
}

func NewLc3Enc() (*GstLc3Enc, error) {
	e, err := gst.NewElement("lc3enc")
	if err != nil {
		return nil, err
	}
	return &GstLc3Enc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewLc3EncWithName(name string) (*GstLc3Enc, error) {
	e, err := gst.NewElementWithName("lc3enc", name)
	if err != nil {
		return nil, err
	}
	return &GstLc3Enc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}
