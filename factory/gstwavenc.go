// source: gst-plugins-good

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Encode raw audio into WAV
type GstWavEnc struct {
	*gst.Element
}

func NewWavEnc() (*GstWavEnc, error) {
	e, err := gst.NewElement("wavenc")
	if err != nil {
		return nil, err
	}
	return &GstWavEnc{Element: e}, nil
}

func NewWavEncWithName(name string) (*GstWavEnc, error) {
	e, err := gst.NewElementWithName("wavenc", name)
	if err != nil {
		return nil, err
	}
	return &GstWavEnc{Element: e}, nil
}
