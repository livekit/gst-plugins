// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Adaptive Multi-Rate Wideband audio decoder
type GstAmrwbDec struct {
	*GstAudioDecoder
}

func NewAmrwbDec() (*GstAmrwbDec, error) {
	e, err := gst.NewElement("amrwbdec")
	if err != nil {
		return nil, err
	}
	return &GstAmrwbDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewAmrwbDecWithName(name string) (*GstAmrwbDec, error) {
	e, err := gst.NewElementWithName("amrwbdec", name)
	if err != nil {
		return nil, err
	}
	return &GstAmrwbDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}
